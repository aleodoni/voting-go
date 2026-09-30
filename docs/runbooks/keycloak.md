# Runbook — Keycloak

- Versão nos composes: 24.0 (dev) / 24.0.5 (staging).
- Realm importado na subida (`--import-realm`) a partir de `infra/keycloak/realm-import/`: `voting-realm.json` (dev) e `voting-realm.staging.json` (staging).
- Clients no realm de dev: `voting-web` (front) e `voting-api` (audience validada pela API — `KEYCLOAK_CLIENT_ID`).
- Grupo administrador configurável por `ADMIN_GROUP` (padrão `/admin`).

## Variáveis relacionadas

| Onde | Variáveis |
| ---- | --------- |
| API | `KEYCLOAK_ISSUER`, `KEYCLOAK_CLIENT_ID`, `JWKSURL` |
| Front | `VITE_KEYCLOAK_URL`, `VITE_KEYCLOAK_REALM`, `VITE_KEYCLOAK_CLIENT_ID` |

## Tarefas comuns

- **Token para testes locais**: `scripts/get-token.sh`.
- **Novo usuário**: criar no Keycloak; a credencial (`ativo`, `pode votar`, `pode administrar`) é gerenciada na aplicação (tela *Gerenciar usuários* do admin).
- **Alterar realm**: exportar do Keycloak, revisar e commitar o JSON — **sem segredos**.

## Senhas dos usuários

- **Dev** (`voting-realm.json`): os usuários de teste têm senhas triviais de propósito, para uso **somente local**. Não exponha esse Keycloak fora da máquina de desenvolvimento.
- **Staging** (`voting-realm.staging.json`): o realm **não traz senhas**, e tem proteção contra força bruta e política de senha (`length(12)`). Depois do primeiro import, defina a senha de cada usuário no console (*Users → Credentials*).
- O `--import-realm` só cria o realm se ele ainda **não existir**. Alterar o JSON não muda um Keycloak já em execução: o de staging atual mantém as senhas de teste antigas, por decisão registrada com a infra.