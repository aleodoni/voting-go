# Prompt — nova feature na API

```
Feature: <nome e objetivo>

Contexto de negócio: <regras, quem pode executar (admin? vereador?), eventos SSE esperados>

Implemente de dentro para fora, entregando cada arquivo completo:
1. Migration (.up/.down) se houver mudança de schema — sem editar migrations existentes.
2. Domain: entidade/agregado, erros, evento de domínio (se aplicável), método na interface de repositório.
3. Query em sqlc/queries/*.sql (não editar sqlc/generated; me diga para rodar `sqlc generate`).
4. Repositório (persistence/*_repository_sqlc.go) e mapper.
5. Caso de uso em internal/application/<contexto>/<verbo>_<objeto>.go com Input próprio,
   shared.VerificarAdmin quando for de admin, e o _test.go com fakes atualizados.
6. Handler com anotações swaggo (request/response/mapper) e registro em router/ + bootstrap/.
7. Front, se necessário: hook + componente (use o prompt feature-frontend.md).

No fim: resumo dos arquivos alterados, comandos que devo rodar (make migrate, sqlc generate,
make swagger, make test), checklist de docs a atualizar (antes do commit) e mensagem de commit em inglês num único bloco.
```