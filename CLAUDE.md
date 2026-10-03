# Voting — orientação para o Claude

Leia isto antes de começar a mexer no projeto. Para arquitetura e domínio,
veja `docs/architecture.md` e `docs/domain.md`. Para o que falta fazer, veja
`docs/roadmap.md` (features, com a seção "Próximo passo") e `BACKLOG.md`
(dívida técnica). O prompt de início de sessão está em `docs/prompt.md`.

## O que é

Sistema de votação legislativa: API Go (Gin) + dois apps React/Vite em
monorepo pnpm (`web/apps/admin`, `web/apps/vereador`, `web/packages/shared`).
Clean Architecture/DDD leve em `internal/`, contextos: `votacao`, `usuario`,
`sincronia`, `relatorio`, `job`. Reuniões, projetos e pareceres vêm do banco
SPL via `postgres_fdw` + sincronia; eventos de votação chegam ao front por
SSE.

## Stack

- **API**: Go 1.25, Gin, PostgreSQL (pgx), sqlc, Keycloak (JWT via JWKS), swaggo, go-pdf/fpdf
- **Front** (`web/`): React, Vite, TanStack Router/Query, react-hook-form + zod, shadcn/ui + Tailwind, Biome
- **Infra**: Docker Compose, Keycloak 24 (realm import), GitHub Actions, GHCR/ECR
- **Banco**: migrations em `migrations/` (golang-migrate); FDW em `fdw/`

## Estado atual

- Fluxo de branches: `develop` e `main` (push e PR) rodam o `ci.yml`, com o job `test` (Swagger + testes Go) e o job `web` (Biome + build do front); `main` e `staging` disparam os builds de imagem (`build_api|admin|vereador`); `docs.yml` publica Swagger UI + casos de uso no GitHub Pages a partir da `main`.
- O repositório foi reestruturado: o Swagger gerado agora fica em `swagger/` (ignorado pelo git) e `docs/` é documentação versionada (ADR 0004). O import em `internal/router/router.go` é `github.com/aleodoni/voting-go/swagger`.
- `infra/.env.*` reais não são mais versionados; use os `.example`.
- O realm de staging (`voting-realm.staging.json`) não traz senhas de usuário; o de dev tem senhas triviais só para uso local. O Keycloak de staging **em execução** mantém as senhas de teste antigas, por decisão do Alexandre com a infra: não é pendência, não reabrir (`docs/runbooks/keycloak.md`).
- O front tem um só linter, o Biome, rodado pela raiz de `web/` (`pnpm lint`, limpo hoje) e também pelo job `web` do CI.
- Existe ambiente de **staging** em Docker Compose (`infra/docker-compose.staging.yml`: db, keycloak, api, admin, vereador). Há mais de um caminho de deploy no repo (compose via SSH, AWS/ECR); qual é o oficial ainda não foi definido — ver `docs/runbooks/deploy.md` e `BACKLOG.md`.
- O event bus do SSE é em memória (por instância); escalar a API para mais de uma réplica exige broker (ADR 0003, `docs/roadmap.md`).

## Convenções de código

- Idioma do domínio e do código de negócio: **português** (votacao, reuniao, projeto, parecer, credencial). Comentários de pacote em português.
- Casos de uso: um arquivo por caso de uso (`<verbo>_<objeto>.go`), tipo `XxxUseCase`, `NewXxxUseCase`, `Execute(ctx, input)` com `Input` próprio. Autorização de admin via `shared.VerificarAdmin`.
- Cada caso de uso tem `_test.go` usando fakes de `internal/test/fakes`.
- Handlers: um arquivo por endpoint, com anotações swaggo; request/response/mapper no pacote do handler.
- Handlers traduzem o erro do caso de uso com `httperr.Respond` (`internal/handler/httperr`): 404 para não encontrado, 403 para sem permissão, 500 genérico com o erro real no log. Nunca devolver `err.Error()` direto ao cliente. Os handlers de `votacao` ainda não migraram (`BACKLOG.md`). Handlers que respondem 202 e executam em goroutine devem checar a permissão **antes** de responder.
- SQL novo em `internal/infrastructure/persistence/sqlc/queries/*.sql` + `sqlc generate`. Nunca editar `sqlc/generated/`.
- Migrations: sempre par `.up.sql`/`.down.sql`, numeração sequencial (`make migrate-create name=...`). Nunca editar migration já aplicada.
- Front: Biome (tabs, aspas simples) é o único linter/formatter (ADR 0006): `make lint-web` verifica e `cd web && pnpm lint:fix` corrige. JSON e CSS não são formatados. Código usado pelos dois apps vai para `@voting/shared`.
- Não versionar segredos.

## Comandos úteis

```bash
make bootstrap       # migrate + seed + fdw
make run | dev       # API (dev = air)
make swagger         # gera ./swagger (NÃO usa docs/)
make test            # testes Go
make dev-web         # front (admin :5173, vereador :5174)
make build-web | lint-web
```

## Convenções de trabalho

- **Alexandre aplica as mudanças localmente e testa/builda ele mesmo.**
  Não peça para rodar `go build`/testes por você — peça para ele rodar e
  reportar o resultado.
- **Nunca escrever no GitHub — sem exceção.** Nem commit, nem push, nem
  mesmo se ele parecer estar pedindo isso na mensagem ("commit e push?",
  "pode commitar") — isso é uma pergunta pra ele decidir, não uma
  autorização. Entregar sempre o arquivo completo (não diff) pra ele
  aplicar, commitar e dar push por conta própria. Na dúvida, a resposta é
  não.
- **Depois de qualquer alteração de código/config, verificar se algum
  arquivo em `docs/` precisa ser atualizado em conjunto** (`roadmap.md`,
  `BACKLOG.md`, `domain.md`, runbooks, ADRs). Itens resolvidos são
  **removidos** do `BACKLOG.md`/`roadmap.md` — o commit que resolveu é o
  registro. Entregar os docs atualizados junto com a mudança, não só
  quando perguntado.
- **No fim de cada iteração, ANTES da mensagem de commit, mostrar o
  checklist de documentação**: para cada arquivo (`roadmap.md`,
  `BACKLOG.md`, `CLAUDE.md` "Estado atual", `architecture.md`,
  `domain.md`, ADRs, runbooks, README) dizer "atualizar" ou "sem mudança"
  e o motivo, com base nos passos executados. Entregar os docs
  atualizados junto com a mudança. Isso inclui reescrever a seção
  "Próximo passo" do `docs/roadmap.md`, que é de onde a próxima sessão
  começa. Só depois vem a mensagem de commit.
- **No início da sessão** (prompt de `docs/prompt.md`): ler `CLAUDE.md`,
  `docs/roadmap.md` e `BACKLOG.md`, resumir o estado, dizer qual é o
  próximo passo e começar por ele, salvo indicação contrária.
- **Prefira o arquivo completo a diff/patch** quando o arquivo for pequeno
  ou já tiver sido editado manualmente. Se uma edição pontual não aplicar,
  não insista adivinhando: peça o conteúdo atual (`cat arquivo`) e gere em
  cima dele.
- **Mensagens de commit em inglês**, indicativo, conventional commits
  (`fix(votacao): ...`, `feat(web): ...`, `chore: ...`), com corpo
  explicando o "porquê", não só o "o quê". São sugestão pronta pra ele
  copiar. **Entregar sempre a mensagem inteira num único bloco de código**
  (assunto, linha em branco e corpo), pronta para colar no editor do
  `git commit` — não quebrada em vários `-m`. Uma mudança que cruza back e
  front vai num único commit.
- **Não tente compilar/testar bootstrapando outro toolchain** quando a
  versão do `go.mod` (hoje 1.25) não estiver disponível no sandbox. Ele
  prefere rodar `go build`/`go test` localmente para economizar recursos.
  Para código só com stdlib, `gofmt` basta; se esbarrar em dependência,
  pare e entregue para ele testar.
- **O clone da sessão fica desatualizado** entre mensagens, já que ele
  aplica e dá push por conta própria. Antes de editar algo que pode ter
  mudado, rode `git fetch` e sincronize/compare antes de assumir o estado
  do arquivo local.
- **O token do GitHub, quando fornecido, é somente leitura** na prática.
  Não tente escrever por API nem por `git push`; peça para ele checar o
  status das Actions na aba do GitHub.
- **Antes de decidir entre manter/promover/aposentar duas soluções em
  paralelo, gere uma consulta SQL comparativa** para ele rodar localmente
  e colar o resultado, em vez de decidir só pela leitura do código.
- **Ao propor melhorias de qualidade de dado, não priorizar regras que
  excluem informação** (tratar exceção/outlier por exclusão). Prefira
  abordagens que mantêm todo o dado e reduzem a influência de pontos
  discrepantes de forma proporcional.
- Domínio de votação (`docs/domain.md`): releia antes de mexer em status,
  opções de voto ou credenciais — os termos têm significado específico e
  alguns itens ainda estão marcados como "confirmar".

## Dívida técnica conhecida

Ver `BACKLOG.md`. Não há P0 nem P1 de segurança em aberto; os P1 restantes são de testes (`tests/api`) e de deploy.