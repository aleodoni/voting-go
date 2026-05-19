# ---------- Build Stage ----------
FROM golang:1.26-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./
RUN go mod download

# Instala swag
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copia código
COPY . .

# Gera docs
RUN swag init -g cmd/api/main.go --output docs

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

WORKDIR /app

RUN apk add --no-cache ca-certificates curl

COPY --from=builder /app/api .
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/seeds ./seeds

EXPOSE 8080

ENTRYPOINT ["/app/api"]