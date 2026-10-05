# Roadmap

## Próximo passo

> Mantida pelo Claude ao final de cada iteração: é daqui que a próxima sessão parte.

1. Definir o caminho de deploy oficial (compose/SSH × AWS/ECR) e desativar o que não for usado (`docs/runbooks/deploy.md`). Depende de decisão do Alexandre.
2. Escolher o próximo P2 do `BACKLOG.md`. O de mais retorno agora é o do front: mostrar a mensagem do servidor nos toasts e no `VotingCard` (o 422 já traz, por exemplo, "usuário já votou nesta votação"), em vez de "Request failed with status code 422".

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