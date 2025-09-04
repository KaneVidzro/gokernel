# ========================
# Variables
# ========================
DOMAIN ?= $(shell grep DOMAIN .env | cut -d '=' -f2)
EMAIL  ?= $(shell grep SERVER_EMAIL .env | cut -d '=' -f2)

DOCKER_COMPOSE = docker compose -f docker/docker-compose.yml
MAILU_COMPOSE = docker compose -f docker/mailu/docker-compose.yml

# ========================
# Docker / Services
# ========================

.PHONY: up
up: ## Start all services including Mailu
	$(DOCKER_COMPOSE) up -d --build
	$(MAILU_COMPOSE) up -d


.PHONY: down
down: ## Stop and remove all services including Mailu
	$(DOCKER_COMPOSE) down
	$(MAILU_COMPOSE) down

.PHONY: restart
restart: down up ## Restart all services including Mailu

.PHONY: logs
logs: ## Tail logs from all containers including Mailu
	$(DOCKER_COMPOSE) logs -f &
	$(MAILU_COMPOSE) logs -f

.PHONY: ps
ps: ## Show running containers
	$(DOCKER_COMPOSE) ps

# ========================
# Database
# ========================

.PHONY: db-shell
db-shell: ## Open a psql shell inside the Postgres container
	$(DOCKER_COMPOSE) exec db psql -U $${POSTGRES_USER} -d $${POSTGRES_DB}

.PHONY: db-migrate
db-migrate: ## Run database migrations
	go run ./cmd/migrate

# ========================
# gRPC / Protobuf
# ========================

PROTO_SRC = proto
PROTO_OUT = proto

.PHONY: proto
proto: ## Generate gRPC + protobuf Go code
	protoc -I=$(PROTO_SRC) \
	    --go_out=$(PROTO_OUT) --go_opt=paths=source_relative \
	    --go-grpc_out=$(PROTO_OUT) --go-grpc_opt=paths=source_relative \
	    $(PROTO_SRC)/*.proto

# ========================
# SSL / Certbot
# ========================

.PHONY: cert-init
cert-init: ## One-time: request initial Let's Encrypt certificate
	$(DOCKER_COMPOSE) run --rm certbot certonly --webroot \
	    -w /var/www/html \
	    -d $(DOMAIN) --email $(EMAIL) --agree-tos --no-eff-email

.PHONY: cert-renew
cert-renew: ## Force renew all certificates
	$(DOCKER_COMPOSE) run --rm certbot renew --webroot -w /var/www/html --quiet

# ========================
# Go Build
# ========================

.PHONY: build-api
build-api: ## Build API binary
	go build -o bin/api ./cmd/api

.PHONY: build-grpc
build-grpc: ## Build gRPC binary
	go build -o bin/grpc ./cmd/grpc

.PHONY: clean
clean: ## Clean build artifacts
	rm -rf bin/

# ========================
# Quality: Linting & Tests
# ========================

.PHONY: lint
lint: ## Run golangci-lint (requires install)
	golangci-lint run ./...

.PHONY: test
test: ## Run all Go tests with coverage
	go test ./... -cover

# ========================
# Mailu
# ========================

.PHONY: mailu-up
mailu-up: ## Start only Mailu services
	$(MAILU_COMPOSE) up -d

.PHONY: mailu-down
mailu-down: ## Stop only Mailu services
	$(MAILU_COMPOSE) down
