# ---------- Build Stage ----------
FROM golang:1.26-alpine AS builder

# Diretório de trabalho dentro do container
WORKDIR /app

# Dependências do sistema para build
RUN apk add --no-cache git ca-certificates

# Copia dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia código
COPY . .

# Argumentos para cross-compiling
ARG TARGETOS
ARG TARGETARCH

# Compila API enxuta
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build \
    -ldflags="-s -w" \
    -trimpath \
    -o api \
    ./cmd/api

# ---------- Runtime Stage ----------
FROM alpine:3.20

# Diretório de trabalho
WORKDIR /app

# Apenas runtime essencial
RUN apk add --no-cache ca-certificates curl

# Copia binário
COPY --from=builder /app/api .

# Copia assets
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/seeds ./seeds

# Porta exposta
EXPOSE 8080

# Executa API
ENTRYPOINT ["/app/api"]