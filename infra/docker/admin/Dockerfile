# === Stage 0: Builder ===
FROM node:20-alpine AS builder

WORKDIR /app

RUN corepack enable && corepack prepare pnpm@10.33.0 --activate

COPY package.json ./package.json
COPY pnpm-lock.yaml ./pnpm-lock.yaml
COPY pnpm-workspace.yaml ./pnpm-workspace.yaml

COPY apps ./apps
COPY packages ./packages

RUN pnpm install --no-frozen-lockfile

RUN pnpm --filter @voting/admin build

# === Stage 1: Nginx ===
FROM nginx:alpine

COPY --from=builder /app/apps/admin/dist /usr/share/nginx/html
COPY --from=builder /app/apps/admin/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]