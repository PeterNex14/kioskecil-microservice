# kioskecil_microservice Makefile

# Default environment file
ENV_FILE ?= .env.development

# Load env file to shell for current make process (optional, but good for local checks)
include $(ENV_FILE)
export $(shell sed 's/=.*//' $(ENV_FILE))

.PHONY: up down restart logs logs-f migrate-status migrate-new db-shell generate tidy build help

# --- Docker Lifecycle ---

up:
	@echo "Starting containers with $(ENV_FILE)..."
	docker compose --env-file $(ENV_FILE) up -d

down:
	@echo "Stopping containers..."
	docker compose --env-file $(ENV_FILE) down

restart: down up

logs:
	docker compose --env-file $(ENV_FILE) logs $(s)

logs-f:
	docker compose --env-file $(ENV_FILE) logs -f $(s)

build:
	@echo "Rebuilding images..."
	docker compose --env-file $(ENV_FILE) build

# --- Database Migrations (Embedded & Status) ---

migrate-status:
	@echo "Checking migration status in $(USER_DB_NAME)..."
	docker compose --env-file $(ENV_FILE) exec db_kios psql -U $(USER_DB_USER) -d $(USER_DB_NAME) -c "SELECT version_id, is_applied, tstamp FROM goose_db_version ORDER BY id DESC;"

migrate-new:
	@if [ -z "$(NAME)" ]; then echo "Error: NAME is required. Usage: make migrate-new NAME=migration_name"; exit 1; fi
	@TIMESTAMP=$$(date +%Y%m%d%H%M%S); \
	FILE="user-service/db/migrations/$${TIMESTAMP}_$(NAME).sql"; \
	printf -- "-- +goose Up\n-- +goose StatementBegin\n-- SQL in section 'Up' is executed when this migration is applied\n-- +goose StatementEnd\n\n-- +goose Down\n-- +goose StatementBegin\n-- SQL section 'Down' is executed when this migration is rolled back\n-- +goose StatementEnd\n" > $$FILE; \
	echo "Created migration: $$FILE"

db-shell:
	@echo "Entering database shell ($(USER_DB_NAME))..."
	docker compose --env-file $(ENV_FILE) exec db_kios psql -U $(USER_DB_USER) -d $(USER_DB_NAME)

# --- Code Generation (SQLC) ---

generate:
	@echo "Generating code from SQL..."
	docker run --rm -v $(shell pwd):/src -w /src/user-service sqlc/sqlc generate

# --- Go Utilities ---

# Run go mod tidy in Docker to ensure toolchain consistency
tidy:
	@echo "Cleaning up Go modules (common)..."
	docker run --rm -v $(shell pwd):/app -w /app/common golang:1.25-alpine go mod tidy
	@echo "Cleaning up Go modules (user-service)..."
	docker run --rm -v $(shell pwd):/app -w /app/user-service golang:1.25-alpine go mod tidy

# --- Help ---

help:
	@echo "Available commands:"
	@echo "  up             Start containers (default env: .env.development)"
	@echo "  down           Stop containers"
	@echo "  restart        Restart containers"
	@echo "  logs           View container logs"
	@echo "  logs-f         Follow container logs in real time"
	@echo "  build          Rebuild containers"
	@echo "  migrate-status Show migration status in database"
	@echo "  migrate-new    Create new migration (use NAME=...)"
	@echo "  db-shell       Enter PostgreSQL shell"
	@echo "  generate       Generate code with SQLC (Docker)"
	@echo "  tidy           Run go mod tidy in all services"
