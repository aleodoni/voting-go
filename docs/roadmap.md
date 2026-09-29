# Roadmap

## Próximo passo

> Mantida pelo Claude ao final de cada iteração: é daqui que a próxima sessão parte.

1. **Trocar as senhas dos usuários no Keycloak de staging em execução** (`BACKLOG.md`, P0). É ação do Alexandre no console; o realm versionado já não traz senhas.
2. CI/CD (fase 4): `ci.yml` também em `main`/PRs e um job do front (`pnpm install --frozen-lockfile`, `pnpm lint`, `pnpm build`). O lint (Biome, ADR 0006) já passa limpo e os três pacotes compilam.
3. Definir o caminho de deploy oficial (compose/SSH × AWS/ECR) e desativar o que não for usado (`docs/runbooks/deploy.md`).

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
- [x] Reestruturação do repositório (higiene, `docs/`, Swagger fora de `docs/`, código duplicado do front em `@voting/shared`, Biome como único linter)

## Em andamento

- Nada no momento.

## Planejado

- [ ] Dashboard de votação
- [ ] Auditoria de votos (quem votou o quê, quando, alterações)
- [ ] Suporte a múltiplas réplicas da API (broker para eventos SSE em vez do bus em memória)
- [ ] CI completo (Go + front) e deploy padronizado
- [ ] Testes E2E do fluxo abrir → votar → fechar

## Ideias

- [ ] Publicar `docs/` junto do Swagger no GitHub Pages
- [ ] Observabilidade (métricas de conexões SSE, tempo de votação)