# Arquitetura

## Camadas (API)

```
handler ──► application ──► domain ◄── infrastructure
(HTTP)      (casos de uso)  (regras)    (sqlc, PDF, mappers)
```

| Camada | Pasta | Responsabilidade |
| ------ | ----- | ---------------- |
| Handler | `internal/handler/<contexto>` | HTTP: bind, validação de formato, chamada ao caso de uso, response, anotações Swagger |
| Application | `internal/application/<contexto>` | Casos de uso (`Execute(ctx, input)`), autorização (`shared.VerificarAdmin`), orquestração, transações |
| Domain | `internal/domain/<contexto>` | Entidades, agregados (`go-ddd`), eventos de domínio, interfaces de repositório |
| Infrastructure | `internal/infrastructure` | Repositórios sqlc, mappers, geração de PDF |
| Platform | `internal/platform` | Event bus, JWT, geração de IDs (cuid), transações |
| Bootstrap | `internal/bootstrap` | Composição: repositórios → casos de uso → handlers → router |

Contextos: `votacao` (reunião, projeto, parecer, votação, voto), `usuario` (usuário, credencial), `sincronia`, `relatorio`, `job`.

Regra de dependência: `domain` não importa nenhuma outra camada; `application` depende de interfaces do domínio; `infrastructure` implementa essas interfaces.

## Rotas e autenticação

- `/api/v1/*` protegido por `JWTMiddleware` (issuer, audience e assinatura via JWKS do Keycloak). Exceção: `/health`.
- Na subida, a API carrega o JWKS com até 6 tentativas (espera crescente, cerca de 25s no total); se o Keycloak não responder, ela encerra com erro.
- `/api/v1/eventos` (SSE) é protegido como as demais rotas: `Authorization: Bearer`, sem token na query string (ADR 0007).
- `/internal/jobs/*` usa `InternalJobMiddleware` (token estático `JOBS_TOKEN`, comparação em tempo constante) — feito para chamadas de agendadores, não de usuários.
- Autorização de negócio (admin, pode votar) vem da **credencial** do usuário no banco, verificada nos casos de uso — não apenas do token.
- Erros de caso de uso viram resposta HTTP em `internal/handler/httperr`: 404 (não encontrado), 403 (sem permissão) e 500 genérico, com o erro real no log. Hoje só os handlers `usuario/*` usam; os demais ainda respondem 403 a qualquer erro (`BACKLOG.md`).

## Fluxo de votação

1. Admin abre a votação de um projeto → `Votacao.Abrir` levanta `VotacaoAbertaEvent`.
2. Eventos de domínio são publicados no `event.Bus`; cada cliente com conexão SSE recebe o evento.
3. Vereadores conectados veem a votação aberta (`GET /votacao/aberta` + SSE) e registram voto (`POST /votacao/{id}/voto`), gerando `VotoRegistradoEvent`.
4. Admin acompanha o painel (progresso, lista de votos) e fecha (`VotacaoFechadaEvent`) ou cancela (`VotacaoCanceladaEvent`).
5. Há um job (`/internal/jobs/fecha_abertas`) que fecha votações que ficaram abertas.
6. O relatório PDF da reunião consolida projetos e votações.

Restrições no banco: um voto por usuário por votação (`uq_usuario_voto`), uma votação por projeto (`unique projeto_id`). Existe também a função SQL `f_save_vote` (migration 000024) — confirmar se é o caminho usado pelo repositório ao registrar voto.

## SSE / Event bus

`platform/event.Bus` mantém os assinantes em memória (`map[chan Event]*Subscriber`), com canal bufferizado (10). Uma nova conexão do mesmo usuário substitui a anterior. `GET /usuarios-conectados` (admin) lê essa lista.

Consequência: o estado dos assinantes é **por instância** da API. Escalar para mais de uma réplica exige um broker externo (ver roadmap).

## Integração com o banco SPL

Reuniões, projetos e pareceres nascem no banco SPL. A API acessa esse banco via `postgres_fdw` (`fdw/`, servidor `server_spl`) e uma procedure de sincronização diária (`p_spl_daily_sync`) copia os dados para tabelas locais. Cada execução é registrada na tabela `sincronia` (contadores e status). A sincronização pode ser disparada por um admin (`POST /sincronia`) ou por job interno.

Setup: `make fdw` (aplica `fdw/spl_setup.sql` com as variáveis `DB_SPL_*`). O CLI (`voting-cli fdw|fdw-drop`) faz o mesmo em ambientes de deploy.

## Persistência

- SQL escrito à mão em `sqlc/queries/*.sql`; código gerado em `sqlc/generated/` (não editar).
- Lógica pesada em funções/views PostgreSQL versionadas em migrations (`f_get_projetos_completo`, `f_get_voting_stats`, `v_reunioes_hoje`…).
- Transações via `TxManager`/`UnitOfWork` sqlc.

## Frontend (`web/`)

- Monorepo pnpm: `apps/admin` (:5173), `apps/vereador` (:5174), `packages/shared` (`@voting/shared`).
- TanStack Router (file-based, `routeTree.gen.ts` é gerado), TanStack Query para dados, `useSSE` (em shared) para eventos.
- Autenticação com `keycloak-js` (`AuthContext`, `keycloak.ts` em shared); cliente HTTP em `api-client.ts`.
- UI: shadcn/ui + Tailwind; componentes base em `shared/src/components/ui`.

## Infra

- `infra/docker-compose.yml` (dev: Postgres + Keycloak); `infra/docker-compose.staging.yml` (db, keycloak, api, admin, vereador).
- Um Dockerfile por serviço em `infra/docker/{api,admin,vereador}`.
- Workflows: `build_api|admin|vereador` (imagens no GHCR), `staging`/`production` (migrate, seed/fdw), `deploy_aws`, `docs` (Swagger UI + casos de uso no Pages), `ci`.