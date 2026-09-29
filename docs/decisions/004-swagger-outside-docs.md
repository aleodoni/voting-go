# ADR 0004 — Swagger gerado em `swagger/`, `docs/` versionada

- **Status**: Aceito
- **Contexto**: o `swag init` gerava o pacote em `docs/` (ignorado pelo git), o que impedia versionar documentação nessa pasta.
- **Decisão**: `swag init ... --output swagger`; `swagger/` ignorado; `docs/` passa a ser documentação versionada. O import no router é `github.com/aleodoni/voting-go/swagger`.
- **Consequências**: Makefile, Dockerfile da API e `router.go` precisam apontar para `swagger`. O workflow `docs.yml` já usa saída própria (`_site`) e não é afetado.