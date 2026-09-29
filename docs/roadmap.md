# Roadmap

## Próximo passo

> Mantida pelo Claude ao final de cada iteração: é daqui que a próxima sessão parte.

1. Fechar a reestruturação: confirmar que o CI do `develop` passa e, se os `infra/.env.*` versionados tinham valores reais, rotacionar as credenciais (`BACKLOG.md`, P0).
2. Frontend (fase 3): mover para `@voting/shared` o que está duplicado entre `admin` e `vereador` (`useIsProjectVoting`, `useUser`, `FormUserInfo`, `LoggedUserCard`) e remover dependências não usadas (`react-router-dom`).
3. CI/CD (fase 4): `ci.yml` também em `main`/PRs, job de lint + build do front, e definir o caminho de deploy oficial.

## Concluído

- [x] Autenticação Keycloak (JWT) e credenciais por usuário
- [x] Reuniões, projetos e pareceres via SPL (FDW + sincronia)
- [x] Abertura, fechamento e cancelamento de votações
- [x] Registro de voto com restrição e voto contrário
- [x] Apuração automática de votos
- [x] Votação em tempo real via SSE
- [x] Relatório PDF por reunião
- [x] Frontend admin e vereador (monorepo pnpm)
- [x] Jobs internos (sincronia, fechamento de votações abertas)
- [x] CLI de operação (`voting-cli`)
- [x] Ambiente de staging com Docker Compose

## Em andamento

- [ ] Reestruturação do repositório (higiene, `docs/`, Swagger fora de `docs/`) — ver `BACKLOG.md`

## Planejado

- [ ] Dashboard de votação
- [ ] Auditoria de votos (quem votou o quê, quando, alterações)
- [ ] Suporte a múltiplas réplicas da API (broker para eventos SSE em vez do bus em memória)
- [ ] CI completo (Go + front) e deploy padronizado
- [ ] Testes E2E do fluxo abrir → votar → fechar

## Ideias

- [ ] Publicar `docs/` junto do Swagger no GitHub Pages
- [ ] Observabilidade (métricas de conexões SSE, tempo de votação)