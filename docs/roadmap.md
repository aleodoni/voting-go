# Roadmap

## Próximo passo

> Mantida pelo Claude ao final de cada iteração: é daqui que a próxima sessão parte.

1. Definir o caminho de deploy oficial (compose/SSH × AWS/ECR) e desativar o que não for usado (`docs/runbooks/deploy.md`). Depende de decisão do Alexandre.
2. Testes P1 (`BACKLOG.md`): separar `tests/api/betha-*.test.js` (serviço externo, exige segredos) em `tests/integrations/` e cobrir os demais `tests/api/*.test.js` no `make test-api`, ou documentar por que não.
3. Backend P2, do `BACKLOG.md`: a inconsistência de `Votacao.Abrir` (`VotacaoAbertaEvent` sem `domainEvent`, `OccurredAt` zerado) e os nomes de arquivo com typo.

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