SHELL := /bin/bash
.SHELLFLAGS := -c
# ============================================================
# Voting Makefile — desenvolvimento local apenas
# Para CI/CD use o voting-cli (cmd/cli/main/main.go)
# ============================================================

APP_NAME=voting-api
IMAGE_NAME=aleodoni/voting-api
TAG ?= $(shell git rev-parse --short HEAD)

ENV_FILE=.env

# ============================================================
# Load .env and generate DB_URL
# ============================================================
define load_env
set -a && . ./$(ENV_FILE) && set +a && \
DB_URL="postgres://$$DBUSER:$$DBPASSWORD@$$DBHOST:$$DBPORT/$$DBNAME?sslmode=$${DBSSLMODE:-disable}"
endef

# ============================================================
# DATABASE
# ============================================================

.PHONY: migrate
migrate:
	@echo "→ Running migrations"
	@$(load_env) && migrate -path migrations -database "$$DB_URL" up

.PHONY: migrate-down
migrate-down:
	@$(load_env) && migrate -path migrations -database "$$DB_URL" down 1

.PHONY: migrate-version
migrate-version:
	@$(load_env) && migrate -path migrations -database "$$DB_URL" version

.PHONY: migrate-create
migrate-create:
	@echo "→ Creating migration: $(name)"
	migrate create -ext sql -dir migrations -seq $(name)

.PHONY: migrate-force
migrate-force:
	@$(load_env) && migrate -path migrations -database "$$DB_URL" force $(version)

.PHONY: seed
seed:
	@echo "→ Running seed"
	@$(load_env) && psql "$$DB_URL" -f seeds/seed.sql

.PHONY: fdw
fdw:
	@echo "→ Running FDW setup"
	@$(load_env) && echo "DB_URL: $$DB_URL" && \
	envsubst < fdw/spl_setup.sql | psql "$$DB_URL"

.PHONY: bootstrap
bootstrap: migrate seed fdw

# ============================================================
# APP
# ============================================================

.PHONY: dev
dev:
	air

.PHONY: run
run:
	go run ./cmd/api/main.go

.PHONY: build
build:
	go build -o bin/api ./cmd/api

# ============================================================
# WEB
# ============================================================

.PHONY: dev-web
dev-web:
	cd web && pnpm dev

# ============================================================
# DOCKER LOCAL
# ============================================================

.PHONY: docker-build-local
docker-build-local:
	docker build -t $(IMAGE_NAME):$(TAG) -f infra/docker/api/Dockerfile .

.PHONY: docker-build-amd64
docker-build-amd64:
	docker buildx build --platform linux/amd64 -t $(IMAGE_NAME):$(TAG) -f infra/docker/api/Dockerfile .

.PHONY: docker-build-arm64
docker-build-arm64:
	docker buildx build --platform linux/arm64 -t $(IMAGE_NAME):$(TAG) -f infra/docker/api/Dockerfile .

.PHONY: docker-build-multi
docker-build-multi:
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		-t $(IMAGE_NAME):$(TAG) \
		-f infra/docker/api/Dockerfile .

.PHONY: docker-compose-up
docker-compose-up:
	docker compose -f ./infra/docker-compose.yml up -d

.PHONY: docker-compose-down
docker-compose-down:
	docker compose -f ./infra/docker-compose.yml down

# ============================================================
# TESTS
# ============================================================

.PHONY: test
test:
	gotestsum --format-hide-empty-pkg --format testname ./...

.PHONY: test-health
test-health:
	k6 run tests/api/health.test.js

.PHONY: test-api
test-api: seed test-health

# ============================================================
# SWAGGER
# ============================================================

.PHONY: swagger
swagger:
	swag fmt
	swag init -g cmd/api/main.go --parseInternal

# ============================================================
# CLI
# ============================================================

.PHONY: build-cli
build-cli:
	go build -ldflags="-s -w" -o voting-cli cmd/cli/main/main.go

.PHONY: run-cli
run-cli:
	go run cmd/cli/main/main.go