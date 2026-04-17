# kioskecil_microservice Makefile

# Default environment file
ENV_FILE ?= .env.development

# Load env file to shell for current make process (optional, but good for local checks)
include $(ENV_FILE)
export $(shell sed 's/=.*//' $(ENV_FILE))

.PHONY: up down restart logs migrate-up migrate-down migrate-status migrate-new generate tidy build help

# --- Docker Lifecycle ---

up:
	@echo "Starting containers with $(ENV_FILE)..."
	docker compose --env-file $(ENV_FILE) up -d

down:
	@echo "Stopping containers..."
	docker compose down

restart: down up

logs:
	docker compose logs $(s)

logs-f:
	docker compose logs -f $(s)

build:
	@echo "Rebuilding images..."
	docker compose --env-file $(ENV_FILE) build

# --- Database Migrations (Goose) ---

migrate-up:
	@echo "Running migrations up..."
	docker compose --env-file $(ENV_FILE) run --rm migrate-user up

migrate-down:
	@echo "Running migrations down..."
	docker compose --env-file $(ENV_FILE) run --rm migrate-user down

migrate-status:
	@echo "Checking migration status..."
	docker compose --env-file $(ENV_FILE) run --rm migrate-user status

migrate-new:
	@echo "Creating new migration: $(NAME)..."
	@if [ -z "$(NAME)" ]; then echo "Error: NAME is required. Usage: make migrate-new NAME=migration_name"; exit 1; fi
	goose -dir user-service/db/migrations create $(NAME) sql

db-shell:
	@echo "Entering database shell ($(USER_DB_NAME))..."
	docker compose exec db_kios psql -U $(USER_DB_USER) -d $(USER_DB_NAME)

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
	@echo "  build          Rebuild containers"
	@echo "  migrate-up     Run database migrations up"
	@echo "  migrate-down   Rollback database migrations"
	@echo "  migrate-status Show migration status"
	@echo "  migrate-new    Create new migration (use NAME=...)"
	@echo "  generate       Generate code with SQLC (Docker)"
	@echo "  tidy           Run go mod tidy in all services"
