#!/bin/bash

KEYCLOAK_URL="https://auth.alexandre.odoni.nom.br"
REALM="voting-realm"
CLIENT_ID="voting-web"  # usar client do frontend

USERNAME="$1"
PASSWORD="$2"

if [[ -z "$USERNAME" || -z "$PASSWORD" ]]; then
  echo "Uso: $0 <usuario> <senha>"
  exit 1
fi

# gera token limpo
curl -s -X POST "$KEYCLOAK_URL/realms/$REALM/protocol/openid-connect/token" \
  -d "grant_type=password" \
  -d "client_id=$CLIENT_ID" \
  -d "username=$USERNAME" \
  -d "password=$PASSWORD" \
  | jq -r '.access_token'