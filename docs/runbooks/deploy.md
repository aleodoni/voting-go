# Runbook — Deploy

> Estado atual: há mais de um caminho de deploy no repositório. **Defina qual é o oficial** e marque aqui (item aberto no `BACKLOG.md`).

## Imagens

- Workflows `build_api`, `build_admin`, `build_vereador` publicam no GHCR: `main` → `latest` + `<sha7>`; demais branches → `staging` + `<sha7>-staging`.
- Dockerfiles em `infra/docker/{api,admin,vereador}`. A imagem da API gera o Swagger no build (`--output swagger`).

## Banco (migrate / seed / fdw)

Workflows manuais (`workflow_dispatch`):

- **Deploy Staging**: migrate, seed, fdw (variáveis `STAGING_*` nos secrets).
- **Deploy Production**: migrate, fdw (variáveis `PROD_*`); **sem seed**.

O CLI recusa `seed` fora de `development`/`staging`.

## Stack via Docker Compose (SSH)

`infra/Makefile` (`make build|deploy|logs|rollback TARGET=<orange|staging|production>`) usa `infra/.env.<target>` e `docker-compose.<target>.yml`.

1. Criar `infra/.env.<target>` a partir do `.env.<target>.example` (arquivo real **não** é versionado).
2. `make build TARGET=<target>` e `make deploy TARGET=<target>`.
3. Verificar `GET /api/v1/health` e `health.txt` dos apps.

Rollback: `make rollback TAG=<sha> TARGET=<target>`.

## AWS

Workflow `deploy_aws` (manual, escolhe `production`/`staging`, opção `force_deploy`) — repositório ECR `voting-go`, região `us-east-1`. *Documentar aqui o passo a passo real e se ainda é usado.*

## Checklist pós-deploy

- [ ] `/api/v1/health` responde
- [ ] Login no admin e no vereador (Keycloak). Em um staging novo, defina antes as senhas dos usuários (`docs/runbooks/keycloak.md`)
- [ ] SSE conecta (`/api/v1/eventos`)
- [ ] Última sincronia com sucesso (`GET /sincronia`)