# Voting

Sistema de votação para sessões legislativas: **API em Go** + **frontend React/Vite** (dois apps).
Permite acompanhar reuniões do dia, abrir/fechar votações de projetos, registrar votos em tempo real (SSE) e gerar relatório PDF por reunião.

## Visão geral

| Parte | Tecnologia | Onde |
| ----- | ---------- | ---- |
| API | Go, Gin, PostgreSQL, sqlc, Keycloak (JWT) | `cmd/`, `internal/` |
| Painel administrativo | React, Vite, TanStack Router/Query, shadcn/ui | `web/apps/admin` |
| Painel do vereador | React, Vite, TanStack Router/Query, shadcn/ui | `web/apps/vereador` |
| Código compartilhado do front | pacote `@voting/shared` | `web/packages/shared` |
| Banco / integração | migrations, FDW para o banco SPL | `migrations/`, `fdw/`, `spl/` |
| Infra | Docker Compose, Keycloak realm, Dockerfiles | `infra/` |

Documentação técnica e de processo em [`docs/`](docs/README.md). Contexto para sessões com o Claude em [`CLAUDE.md`](CLAUDE.md).

## Funcionalidades

- Autenticação via Keycloak (JWT) com credenciais próprias por usuário (ativo, pode votar, pode administrar)
- Reuniões do dia, projetos e pareceres sincronizados do banco SPL
- Abertura, fechamento e cancelamento de votações (admin)
- Registro de votos com restrição e voto contrário
- Eventos em tempo real via SSE (votação aberta/fechada/cancelada, voto registrado, usuários conectados)
- Relatório PDF por reunião
- Jobs internos (sincronização e fechamento de votações abertas)
- CLI de operação (`voting-cli`): migrate, seed, fdw

---

## Como rodar localmente

### 1. Pré-requisitos

Go, Docker, [`golang-migrate`](https://github.com/golang-migrate/migrate), [`sqlc`](https://sqlc.dev), [`swag`](https://github.com/swaggo/swag), `psql`, `envsubst`, `pnpm` (front) e opcionalmente `air`, `gotestsum` e `k6`.

### 2. Variáveis de ambiente

```bash
cp .env.example .env
```

| Variável | Descrição | Padrão (dev) |
| -------- | --------- | ------------ |
| `APPNAME` / `APPVERSION` | Nome e versão da API | `Voting API` / `1.0.0` |
| `APPPORT` | Porta da API | `8080` |
| `APPENV` | `development`, `staging` ou `production` | `development` |
| `ALLOW_ORIGINS` | Origens permitidas no CORS (separadas por vírgula) | `http://localhost:5173,http://localhost:5174` |
| `DBHOST` `DBPORT` `DBUSER` `DBPASSWORD` `DBNAME` `DBSSLMODE` | Banco PostgreSQL da aplicação | `localhost` / `15432` / `postgres` / `postgres` / `voting_db` / `disable` |
| `DB_SPL_HOST` `DB_SPL_PORT` `DB_SPL_NAME` `DB_SPL_USER` `DB_SPL_PASSWORD` | Banco SPL (origem de reuniões, projetos e pareceres, via FDW) | — |
| `KEYCLOAK_ISSUER` | URL do realm | `http://localhost:8081/realms/voting-realm` |
| `KEYCLOAK_CLIENT_ID` | Audience esperada no JWT | `voting-api` |
| `JWKSURL` | Endpoint de chaves públicas do Keycloak | `…/protocol/openid-connect/certs` |
| `JOBS_TOKEN` | Token das rotas `/internal/jobs/*` | — |
| `ADMIN_GROUP` | Grupo Keycloak considerado administrador | `/admin` |
| `TEST_USER` / `TEST_PASSWORD` | Usuário do realm de dev usado por `make token` e `make test-api` (só testes) | `usuario.admin` / — |

### 3. Dependências (Postgres + Keycloak)

```bash
make docker-compose-up
```

### 4. Banco de dados

```bash
make bootstrap   # migrate + seed + fdw
```

Comandos individuais: `make migrate`, `make migrate-down`, `make migrate-create name=<nome>`, `make seed`, `make fdw`.

### 5. Swagger

```bash
make swagger
```

O Swagger é gerado em `swagger/` (ignorado pelo git). Rode sempre que alterar as anotações dos handlers. A pasta `docs/` é documentação versionada e **não** recebe arquivos gerados.

### 6. API

```bash
make run    # go run ./cmd/api/main.go
make dev    # com hot reload (air)
```

- API: `http://localhost:8080`
- Swagger UI: `http://localhost:8080/swagger/index.html`

### 7. Frontend

```bash
cd web && pnpm install
make dev-web    # admin em :5173, vereador em :5174
```

Cada app tem um `.env.example` (`VITE_API_URL`, `VITE_KEYCLOAK_URL`, `VITE_KEYCLOAK_REALM`, `VITE_KEYCLOAK_CLIENT_ID`).

---

## Autenticação

A API valida **JWT emitido pelo Keycloak** (issuer, audience e assinatura via JWKS).

```
Authorization: Bearer <token>
```

- No Swagger UI, use **Authorize** e informe `Bearer <token>`.
- O endpoint SSE (`/api/v1/eventos`) usa o mesmo header `Authorization`; o token **não** é aceito por query string. Use um cliente que envie headers (como `fetch`), não o `EventSource` nativo.
- Rotas `/internal/jobs/*` usam um token interno (`JOBS_TOKEN`), não o JWT do usuário.

Para obter um token em desenvolvimento: `scripts/get-token.sh`.

---

## Endpoints

Prefixo `/api/v1`. Detalhes de request/response no Swagger UI.

| Método | Rota | Descrição | Auth |
| ------ | ---- | --------- | ---- |
| `GET` | `/health` | Health check | ❌ |
| `GET` | `/me` | Usuário autenticado | ✅ |
| `GET` | `/usuarios` | Pesquisa usuários (admin) | ✅ |
| `GET` | `/usuarios/{usuarioId}` | Retorna um usuário (admin) | ✅ |
| `PUT` | `/usuarios/fantasia` | Atualiza nome fantasia | ✅ |
| `PUT` | `/usuarios/fantasia-credenciais` | Atualiza nome fantasia e permissões | ✅ |
| `PATCH` | `/usuarios/{id}/credencial` | Atualiza credencial de um usuário | ✅ |
| `GET` | `/usuarios-conectados` | Usuários com conexão SSE ativa | ✅ |
| `GET` | `/reunioes-dia` | Reuniões do dia | ✅ |
| `GET` | `/reunioes/{reuniaoId}/projetos` | Projetos de uma reunião (admin) | ✅ |
| `GET` | `/reunioes/{reuniaoId}/relatorio` | Relatório PDF da reunião | ✅ |
| `GET` | `/projetos/{projetoId}` | Projeto completo (admin) | ✅ |
| `POST` | `/projetos/{projetoId}/votacao/abrir` | Abre uma votação (admin) | ✅ |
| `POST` | `/projetos/{projetoId}/votacao/fechar` | Fecha uma votação (admin) | ✅ |
| `DELETE` | `/projetos/{projetoId}/votacao` | Cancela uma votação (admin) | ✅ |
| `POST` | `/votacao/{votacaoId}/voto` | Registra um voto | ✅ |
| `GET` | `/votacao/aberta` | Projeto com votação aberta | ✅ |
| `GET` | `/votacao/stats` | Estatísticas de votação do dia (admin) | ✅ |
| `GET` | `/sincronia` | Últimas 3 sincronizações (admin) | ✅ |
| `POST` | `/sincronia` | Executa sincronização (admin) | ✅ |
| `GET` | `/eventos` | Stream SSE | ✅ |
| `POST` | `/internal/jobs/sincronia` | Job de sincronização | 🔑 token interno |
| `POST` | `/internal/jobs/fecha_abertas` | Job: fecha votações abertas | 🔑 token interno |

> `/internal/*` fica fora do prefixo `/api/v1`.

---

## Testes

```bash
make test          # testes Go (gotestsum)
make token         # imprime um JWT do realm de dev (usa TEST_USER e TEST_PASSWORD do .env)
make test-api      # seed + testes k6 de leitura da API (requer API e Keycloak rodando)
k6 run -e TOKEN=$(make -s token) tests/api/<arquivo>.test.js
```

O `make test-api` roda só os testes de leitura (`health`, `me` e `retorna-sincronias`). Os que alteram estado (`sincronia`, `atualiza-fantasia-credenciais` e `reunioes-dia`) são manuais por enquanto. A senha do usuário de teste está no realm de dev (`infra/keycloak/realm-import/voting-realm.json`).

---

## Estrutura do projeto

```
cmd/
  api/                 # Entrypoint da API e anotações gerais do Swagger
  cli/                 # voting-cli (migrate, seed, fdw) + TUI

internal/
  application/         # Casos de uso (votacao, usuario, sincronia, relatorio, jobs)
  domain/              # Entidades, agregados, eventos e interfaces de repositório
  infrastructure/      # Persistência (sqlc + repositórios), mappers, geração de PDF
  handler/             # Handlers HTTP, requests/responses e mappers
  middleware/          # JWT, CORS, token de jobs
  router/              # Registro de rotas
  platform/            # Event bus (SSE), JWT, IDs, transações
  bootstrap/           # Composição de dependências
  config/  database/   # Configuração, conexão, migrate, seed, FDW
  test/fakes/          # Fakes de repositório para testes

migrations/            # Migrations SQL (golang-migrate)
fdw/  spl/  seeds/     # Integração com o banco SPL e dados de seed
infra/                 # Docker Compose, Dockerfiles, realm Keycloak
tests/api/             # Testes k6
web/                   # Monorepo pnpm (apps/admin, apps/vereador, packages/shared)
docs/                  # Documentação, roadmap, ADRs, runbooks e prompts
swagger/               # Gerado por `make swagger` (não versionado)
```

Arquitetura baseada em **Clean Architecture / DDD** — ver [`docs/architecture.md`](docs/architecture.md).

---

## Roadmap

Ver [`docs/roadmap.md`](docs/roadmap.md). Dívida técnica em [`BACKLOG.md`](BACKLOG.md).

---

## Licença

MIT