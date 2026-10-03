# Roadmap

## Próximo passo

> Mantida pelo Claude ao final de cada iteração: é daqui que a próxima sessão parte.

1. Definir o caminho de deploy oficial (compose/SSH × AWS/ECR) e desativar o que não for usado (`docs/runbooks/deploy.md`). Depende de decisão do Alexandre.
2. Backend P2: migrar os handlers restantes (`reuniao`, `votacao`, `relatorio`, `jobs`) para `httperr.Respond`, um pacote por vez, e decidir com o Alexandre se o front deve chamar `login()` só em 401 (`BACKLOG.md`).
3. Backend P2: `cmd/cli/main/` com nível extra.

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
- [x] CI completo (Go + front) em `develop` e `main`: `ci.yml` com os jobs `test` e `web`
- [x] Reestruturação do repositório (higiene, `docs/`, Swagger fora de `docs/`, código duplicado do front em `@voting/shared`, Biome como único linter)

## Em andamento

- Nada no momento.

## Planejado

- [ ] Dashboard de votação
- [ ] Auditoria de votos (quem votou o quê, quando, alterações)
- [ ] Suporte a múltiplas réplicas da API (broker para eventos SSE em vez do bus em memória)
- [ ] Deploy padronizado (um caminho oficial documentado)
- [ ] Testes E2E do fluxo abrir → votar → fechar

## Ideias

- [ ] Publicar `docs/` junto do Swagger no GitHub Pages
- [ ] Observabilidade (métricas de conexões SSE, tempo de votação)