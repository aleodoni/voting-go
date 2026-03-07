APP_NAME=voting-api
IMAGE_NAME=aleodoni/voting-api

DB_URL=postgres://postgres:postgres@localhost:15432/voting_db?sslmode=disable

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