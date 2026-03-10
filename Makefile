APP_NAME=voting-api
IMAGE_NAME=aleodoni/voting-api

DB_URL=postgres://postgres:postgres@localhost:15432/voting_db?sslmode=disable

# -------------------------
# Arquivo .env
# -------------------------
ENV_FILE=.env

# -------------------------
# Helper para rodar k6 com .env
# -------------------------
# Usa a variável SCRIPT para indicar qual arquivo k6 rodar
.PHONY: k6-run
k6-run:
	@if [ -f $(CURDIR)/.env ]; then \
		PUBLIC_KEY=$$(grep PUBLIC_KEY $(CURDIR)/.env | cut -d '=' -f2- | tr -d '[:space:]'); \
		USER_ACCESS=$$(grep USER_ACCESS $(CURDIR)/.env | cut -d '=' -f2- | tr -d '[:space:]'); \
		CPF=$$(grep CPF $(CURDIR)/.env | cut -d '=' -f2- | tr -d '[:space:]'); \
		echo "DEBUG PUBLIC_KEY: '$$PUBLIC_KEY'"; \
		echo "DEBUG USER_ACCESS: '$$USER_ACCESS'"; \
		echo "DEBUG CPF: '$$CPF'"; \
		k6 run -e PUBLIC_KEY="$$PUBLIC_KEY" -e USER_ACCESS="$$USER_ACCESS" -e CPF="$$CPF" $(SCRIPT); \
	else \
		k6 run $(SCRIPT); \
	fi

# -------------------------
# Run local
# -------------------------

dev:
	air

run:
	go run ./cmd/api/main.go

build:
	go build -o bin/api ./cmd/api

# -------------------------
# Docker build
# -------------------------

docker-build:
	docker build -t $(IMAGE_NAME):latest .

docker-build-amd64:
	docker buildx build --platform linux/amd64 -t $(IMAGE_NAME):amd64 .

docker-build-arm64:
	docker buildx build --platform linux/arm64 -t $(IMAGE_NAME):arm64 .

docker-build-multi:
	docker buildx build \
	--platform linux/amd64,linux/arm64 \
	-t $(IMAGE_NAME):latest .

# -------------------------
# Docker compose
# -------------------------
docker-compose-up:
	docker compose -f ./infra/docker-compose.yml up -d

docker-compose-down:
	docker compose -f ./infra/docker-compose.yml down


# -------------------------
# Migrations
# -------------------------

migrate-up:
	migrate -path migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" down 1

migrate-version:
	migrate -path migrations -database "$(DB_URL)" version

migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

migrate-force:
	migrate -path migrations -database "$(DB_URL)" force $(version)

# -------------------------
# Tests
# -------------------------
.PHONY: test test-health test-me test-api test-betha test-betha-matricula test-betha-pessoa-fisica

test:
	gotestsum --format-hide-empty-pkg --format testname ./...

get-token-admin:
	@./scripts/get-token.sh usuario.admin 123456

test-me:
	@TOKEN=$$(./scripts/get-token.sh usuario.vereador 123456); \
	k6 run -e TOKEN=$$TOKEN tests/api/me.test.js

test-health:
	k6 run tests/api/health.test.js

test-api: test-health test-me 

test-betha: test-betha-matricula test-betha-pessoa-fisica

test-betha-matricula:
	SCRIPT=tests/api/betha-matricula.test.js make k6-run

test-betha-pessoa-fisica:
	SCRIPT=tests/api/betha-pessoa-fisica.test.js make k6-run
 