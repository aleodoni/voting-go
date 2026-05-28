#!/bin/bash

# KEYCLOAK_URL="http://localhost:8081"
# REALM="voting-realm"
# CLIENT_ID="voting-web"

KEYCLOAK_URL="https://auth.cmc.pr.gov.br"
REALM="cmc"
CLIENT_ID="prod-votacao-web"


USERNAME="$1"
PASSWORD="$2"

if [[ -z "$USERNAME" || -z "$PASSWORD" ]]; then
  echo "Uso: $0 <usuario> <senha>"
  exit 1
fi

RESPONSE=$(curl -s -X POST \
  "$KEYCLOAK_URL/realms/$REALM/protocol/openid-connect/token" \
  -d "grant_type=password" \
  -d "client_id=$CLIENT_ID" \
  -d "username=$USERNAME" \
  -d "password=$PASSWORD")

TOKEN=$(echo "$RESPONSE" | jq -r '.access_token')

if [[ "$TOKEN" == "null" || -z "$TOKEN" ]]; then
  echo "Erro ao obter token"
  echo "$RESPONSE"
  exit 1
fi

echo "$TOKEN"