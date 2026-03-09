#!/bin/bash

KEYCLOAK_URL="http://localhost:8081"
REALM="voting-realm"
CLIENT_ID="voting-api"

USERNAME=$1
PASSWORD=$2

curl -s -X POST \
"$KEYCLOAK_URL/realms/$REALM/protocol/openid-connect/token" \
-H "Content-Type: application/x-www-form-urlencoded" \
-d "username=$USERNAME" \
-d "password=$PASSWORD" \
-d "grant_type=password" \
-d "client_id=$CLIENT_ID" \
| jq -r '.access_token'