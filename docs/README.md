# Documentação

| Documento | Conteúdo |
| --------- | -------- |
| [architecture.md](architecture.md) | Camadas, fluxo de votação, SSE, integração SPL/FDW, autenticação, frontend |
| [domain.md](domain.md) | Glossário e regras de negócio |
| [roadmap.md](roadmap.md) | Features planejadas e concluídas |
| [decisions/](decisions/) | ADRs — decisões de arquitetura e o porquê |
| [runbooks/](runbooks/) | Deploy, Keycloak, sincronização |

## Prompts

[`prompt.md`](prompt.md) é o prompt de **início de toda sessão**: cole-o numa conversa nova e o Claude acessa o repositório, lê `CLAUDE.md`, `roadmap.md` e `BACKLOG.md`, resume o estado e parte do "Próximo passo". Ao fim da iteração ele mostra o checklist de documentação antes do commit.

Prompts de apoio para tarefas específicas (copie, preencha os campos `<...>` e use junto com o `prompt.md`):

| Arquivo | Quando usar |
| ------- | ----------- |
| [prompts/feature-backend.md](prompts/feature-backend.md) | Novo caso de uso / endpoint na API |
| [prompts/feature-frontend.md](prompts/feature-frontend.md) | Nova tela, componente ou hook |
| [prompts/migration.md](prompts/migration.md) | Nova migration ou mudança de schema |
| [prompts/review.md](prompts/review.md) | Revisão de código ou de estrutura |
| [prompts/bugfix.md](prompts/bugfix.md) | Investigação e correção de bug |

Se uma convenção mudar, ajuste o prompt e o [`CLAUDE.md`](../CLAUDE.md) no mesmo commit.

## Na raiz do repositório

- [`CLAUDE.md`](../CLAUDE.md) — contexto de sessão
- [`BACKLOG.md`](../BACKLOG.md) — dívida técnica

> Esta pasta é documentação versionada. O Swagger gerado (`make swagger`) fica em `swagger/`, ignorado pelo git.

## Como manter

- Mudou uma decisão de arquitetura? Novo ADR em `decisions/` (não edite o antigo; marque como *Substituído*).
- Terminou uma feature? Marque no `roadmap.md`. Achou dívida técnica? `BACKLOG.md`. Resolveu um item do backlog? Remova-o — o commit é o registro.
- Mudou o fluxo de deploy? Atualize o runbook correspondente no mesmo commit.