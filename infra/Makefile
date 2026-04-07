# ============================================================
# Voting Deploy Makefile
#
# Uso:
# make build TARGET=orange
# make deploy TARGET=orange
# make logs TARGET=orange
# make logs-svc SVC=api TARGET=orange
# make rollback TAG=xxxx TARGET=orange
# ============================================================

TARGET ?= local
TAG ?= $(shell git rev-parse --short HEAD)

# ENV_FILE := .env.$(TARGET)

DOCKER_USER := aleodoni

ifeq ($(TARGET),orange)
	PLATFORM := linux/arm64
	ENV_FILE := .env.orange
	COMPOSE_FILE := docker-compose.orange.yml
else
	PLATFORM := linux/amd64
	ENV_FILE := .env.staging
	COMPOSE_FILE := docker-compose.staging.yml
endif

API_IMAGE := $(DOCKER_USER)/voting-api:$(TAG)
ADMIN_IMAGE := $(DOCKER_USER)/voting-admin:$(TAG)
VEREADOR_IMAGE := $(DOCKER_USER)/voting-vereador:$(TAG)
DB_IMAGE := $(DOCKER_USER)/voting-postgres:$(TAG)

HOST     := $(shell grep -E '^SSH_HOST=' $(ENV_FILE) | cut -d= -f2)
SSH_USER := $(shell grep -E '^SSH_USER=' $(ENV_FILE) | cut -d= -f2)
SSH_KEY  := $(shell grep -E '^SSH_KEY='  $(ENV_FILE) | cut -d= -f2)

REMOTE   := $(SSH_USER)@$(HOST)
SSH      := ssh -i $(SSH_KEY) -o StrictHostKeyChecking=no $(REMOTE)
SCP      := scp -i $(SSH_KEY) -o StrictHostKeyChecking=no

REMOTE_DIR := ~/staging

# ────────────────────────────────────────────────────────────
# FRONTEND VARS (leitura do .env)
# ────────────────────────────────────────────────────────────
VITE_API_URL := $(shell grep -E '^VITE_API_URL=' $(ENV_FILE) | cut -d= -f2-)
VITE_KEYCLOAK_URL := $(shell grep -E '^VITE_KEYCLOAK_URL=' $(ENV_FILE) | cut -d= -f2-)
VITE_KEYCLOAK_REALM := $(shell grep -E '^VITE_KEYCLOAK_REALM=' $(ENV_FILE) | cut -d= -f2-)
VITE_KEYCLOAK_CLIENT_ID := $(shell grep -E '^VITE_KEYCLOAK_CLIENT_ID=' $(ENV_FILE) | cut -d= -f2-)

# ────────────────────────────────────────────────────────────
# CHECK ENV
# ────────────────────────────────────────────────────────────
.PHONY: check-env
check-env:
	@test -f $(ENV_FILE) || (echo "ERRO: $(ENV_FILE) não encontrado." && exit 1)
	@test -n "$(HOST)" || (echo "ERRO: SSH_HOST não definido em $(ENV_FILE)" && exit 1)

# ────────────────────────────────────────────────────────────
# BUILD MULTIARCH + PUSH
# ────────────────────────────────────────────────────────────
.PHONY: build
build:
	@echo "→ Build API"
	cd .. && docker buildx build \
		--platform $(PLATFORM) \
		-t $(API_IMAGE) \
		-f infra/docker/api/Dockerfile \
		. --push

	@echo "→ Build Admin"
	cd ../web && docker buildx build \
		--platform $(PLATFORM) \
		--build-arg VITE_API_URL=$(VITE_API_URL) \
		--build-arg VITE_KEYCLOAK_URL=$(VITE_KEYCLOAK_URL) \
		--build-arg VITE_KEYCLOAK_REALM=$(VITE_KEYCLOAK_REALM) \
		--build-arg VITE_KEYCLOAK_CLIENT_ID=$(VITE_KEYCLOAK_CLIENT_ID) \
		-t $(ADMIN_IMAGE) \
		-f ../infra/docker/admin/Dockerfile \
		. --push

	@echo "→ Build Vereador"
	cd ../web && docker buildx build \
		--platform $(PLATFORM) \
		--build-arg VITE_API_URL=$(VITE_API_URL) \
		--build-arg VITE_KEYCLOAK_URL=$(VITE_KEYCLOAK_URL) \
		--build-arg VITE_KEYCLOAK_REALM=$(VITE_KEYCLOAK_REALM) \
		--build-arg VITE_KEYCLOAK_CLIENT_ID=$(VITE_KEYCLOAK_CLIENT_ID) \
		-t $(VEREADOR_IMAGE) \
		-f ../infra/docker/vereador/Dockerfile \
		. --push

	@echo "→ Build Postgres"
	docker buildx build \
		--platform $(PLATFORM) \
		-t $(DB_IMAGE) \
		-f docker/postgres/Dockerfile \
		. --push

# ────────────────────────────────────────────────────────────
# PREPARE ENV
# ────────────────────────────────────────────────────────────
.PHONY: prepare-env
prepare-env: check-env
	cp $(ENV_FILE) .env.deploy && echo "" >> .env.deploy
	echo "API_IMAGE=$(API_IMAGE)" >> .env.deploy
	echo "ADMIN_IMAGE=$(ADMIN_IMAGE)" >> .env.deploy
	echo "VEREADOR_IMAGE=$(VEREADOR_IMAGE)" >> .env.deploy
	echo "DB_IMAGE=$(DB_IMAGE)" >> .env.deploy

# ────────────────────────────────────────────────────────────
# SYNC
# ────────────────────────────────────────────────────────────
.PHONY: sync
sync: prepare-env
	@echo "→ Sincronizando arquivos com $(REMOTE)"
	$(SSH) "mkdir -p $(REMOTE_DIR)/keycloak"
	$(SCP) $(COMPOSE_FILE) $(REMOTE):$(REMOTE_DIR)/docker-compose.yml
	$(SCP) .env.deploy $(REMOTE):$(REMOTE_DIR)/.env
	$(SCP) -r keycloak $(REMOTE):$(REMOTE_DIR)/

# ────────────────────────────────────────────────────────────
# DEPLOY
# ────────────────────────────────────────────────────────────
.PHONY: deploy
deploy: sync
	@echo "→ Deploy em $(HOST)"
	$(SSH) "cd $(REMOTE_DIR) && docker compose pull && docker compose up -d --remove-orphans"
	@echo "✓ Deploy concluído"

# ────────────────────────────────────────────────────────────
# DATABASE (delegado para raiz)
# ────────────────────────────────────────────────────────────
.PHONY: migrate
migrate:
	$(MAKE) -C .. migrate

.PHONY: seed
seed:
	$(MAKE) -C .. seed

.PHONY: fdw
fdw:
	$(MAKE) -C .. fdw

.PHONY: bootstrap
bootstrap:
	$(MAKE) -C .. bootstrap