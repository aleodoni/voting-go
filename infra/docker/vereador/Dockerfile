# === Stage 0: Builder ===
FROM node:20-alpine AS builder

WORKDIR /app

# Ativa o pnpm. A versão vem do campo "packageManager" do package.json.
RUN corepack enable

# Copia arquivos de lock, workspace e a configuração do pnpm (.npmrc).
# Sem o .npmrc o pnpm usa outras configurações e reescreve o lockfile.
COPY package.json ./package.json
COPY .npmrc ./.npmrc
COPY pnpm-lock.yaml ./pnpm-lock.yaml
COPY pnpm-workspace.yaml ./pnpm-workspace.yaml

# Copia código
COPY apps ./apps
COPY packages ./packages

# Instala dependências
RUN pnpm install --frozen-lockfile

# --- Build do Vite com variáveis de ambiente ---
ARG VITE_API_URL
ARG VITE_KEYCLOAK_URL
ARG VITE_KEYCLOAK_REALM
ARG VITE_KEYCLOAK_CLIENT_ID

# Cria .env.local para Vite
RUN echo "VITE_API_URL=${VITE_API_URL}" > apps/vereador/.env.local && \
    echo "VITE_KEYCLOAK_URL=${VITE_KEYCLOAK_URL}" >> apps/vereador/.env.local && \
    echo "VITE_KEYCLOAK_REALM=${VITE_KEYCLOAK_REALM}" >> apps/vereador/.env.local && \
    echo "VITE_KEYCLOAK_CLIENT_ID=${VITE_KEYCLOAK_CLIENT_ID}" >> apps/vereador/.env.local

RUN pnpm --filter @voting/vereador build

# === Stage 1: Nginx ===
FROM nginx:alpine

RUN apk add --no-cache curl

COPY --from=builder /app/apps/vereador/dist /usr/share/nginx/html
COPY --from=builder /app/apps/vereador/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 3001
CMD ["nginx", "-g", "daemon off;"]