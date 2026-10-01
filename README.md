# 🏪 Kioskecil Microservice

A modular, scalable, and type-safe microservice platform designed for warung (small retail shops) and POS (Point-of-Sale) management, built with **Go (Golang 1.25)**, **PostgreSQL 15**, **SQLC**, and **Docker Compose**.

---

## 🛠 Tech Stack & Architecture Highlights

* **Language**: Go (Golang 1.25) utilizing multi-module **Go Workspaces** (`go.work`).
* **Database**: PostgreSQL 15 following the **Database-per-Service** isolation pattern.
* **HTTP Framework**: [Gin Web Framework](https://github.com/gin-gonic/gin).
* **Database Migrations**: [Goose](https://github.com/pressly/goose) with **Embedded Migrations** (`//go:embed`). No separate migration containers required!
* **Type-Safe SQL**: [SQLC](https://sqlc.dev/) for compile-time verified database queries.
* **Structured Logging**: Go's native `log/slog` via `common/logger`.
* **Container Orchestration**: Docker Compose with health checks and zero-downtime dependency chaining.

---

## 📋 Prerequisites

Ensure you have installed on your system:
* [Docker Desktop](https://www.docker.com/products/docker-desktop/) (Docker 20.10+ & Docker Compose v2+)
* `make` (standard on macOS/Linux; for Windows use WSL2 or Git Bash)
* *Optional*: `curl` or [Postman](https://www.postman.com/) to test API endpoints.
*(Note: You do not even need Go or PostgreSQL installed locally on your host machine—everything builds and runs inside Docker!)*

---

## 🚀 Quick Start Guide

Follow these simple steps to get the entire microservice ecosystem running in under 2 minutes:

### 1. Clone the Repository
```bash
git clone https://github.com/PeterNex14/kioskecil-microservice.git
cd kioskecil-microservice
```

### 2. Configure Environment Variables
Copy the example environment configuration:
```bash
cp .env.example .env.development
```

> [!TIP]
> **Port Conflict Note (`HOST_DB_PORT`)**:
> By default, `HOST_DB_PORT=5433` is configured in `.env.development` so it does not collide with any existing PostgreSQL instances running on port `5432` on your machine.
> The microservices inside the Docker network communicate internally on port `5432`.

### 3. Start All Services
Run:
```bash
make up
```

**What happens behind the scenes:**
1. Docker builds the `user-service` image using multi-stage builds.
2. The dedicated PostgreSQL container (`db_users_container`) starts and automatically creates the `db_users` database and user credentials.
3. Once the database passes its health check, `user_service_container` starts up, **automatically executes embedded database migrations** in ~10ms, and begins serving HTTP requests on port `8080`.

---

## 🧪 Verifying the Application

### 1. Check Container Status
```bash
docker compose --env-file .env.development ps
```
You should see:
* `db_users_container`: `Up (healthy)`
* `user_service_container`: `Up`

### 2. Health Check Endpoint
Test that the service is operational:
```bash
curl -i http://localhost:8080/health
```
**Expected Response (200 OK):**
```json
{
  "status": "healthy",
  "message": "User Service is up and running"
}
```

### 3. Register a New User
Test database connectivity and domain logic:
```bash
curl -i -X POST http://localhost:8080/api/v1/register \
  -H "Content-Type: application/json" \
  -d '{
    "id": "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
    "email": "owner@kioskecil.com",
    "password": "securepassword123",
    "full_name": "Pak Budi",
    "role": "owner"
  }'
```
**Expected Response (201 Created):**
```json
{
  "message": "user registered successfully",
  "data": {
    "id": "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
    "email": "owner@kioskecil.com",
    "full_name": "Pak Budi",
    "role": "owner",
    "is_active": true,
    ...
  }
}
```

---

## 🛠 Day-to-Day Developer Workflow

All recurring development tasks are automated via the [`Makefile`](./Makefile):

| Command | Description |
| :--- | :--- |
| `make up` | Start all services in the background (runs pending migrations automatically). |
| `make down` | Stop and remove all running service containers. |
| `make restart` | Rebuild and restart the services. |
| `make logs` | View container logs across all services. |
| `make logs-f` | Stream container logs in real time (`Ctrl+C` to exit). |
| `make build` | Rebuild Docker images after code modifications. |
| `make migrate-status` | Inspect applied database migrations directly from PostgreSQL. |
| `make migrate-new NAME=add_xxx` | Generate a new timestamped SQL migration file in `db/migrations/`. |
| `make db-shell` | Open an interactive `psql` shell into the service database. |
| `make generate` | Re-run SQLC code generator inside Docker. |
| `make tidy` | Run `go mod tidy` cleanly across all modules in Docker. |

---

## 🗄 Database Migrations Guide

Kioskecil uses **Embedded Migrations** via Goose:

### Creating a New Migration
To add a new database table or alter an existing one:
```bash
make migrate-new NAME=create_stores_table
```
This generates a timestamped migration file in `user-service/db/migrations/` (e.g., `20261002_create_stores_table.sql`).

### Applying Migrations
Simply restart the service:
```bash
make restart
# or
make up
```
During startup, Go reads all embedded `.sql` files and applies unapplied migrations instantly.

### Checking Migration Status
```bash
make migrate-status
```

---

## ⚡ Generating Type-Safe SQL (SQLC)

Whenever you add or modify SQL queries in `<service>/db/queries/*.sql`:
```bash
make generate
```
SQLC will automatically compile your queries into type-safe Go structs and functions in `<service>/db/sqlc/`.
*(Note: Never manually edit files inside `db/sqlc/`!)*

---

## 📂 Project Structure

```text
kioskecil-microservice/
├── Makefile                     # Developer task automation
├── docker-compose.yml           # Local multi-service infrastructure
├── go.work                      # Go workspace linking all modules
│
├── common/                      # Shared code across all microservices
│   ├── config/                  # Shared environment helper
│   ├── database/                # Database connector & embedded migration runner
│   ├── logger/                  # Structured slog initialization
│   └── system/                  # OS signal & graceful shutdown handlers
│
├── user-service/                # User & Authentication Microservice
│   ├── Dockerfile               # Multi-stage production-ready Dockerfile
│   ├── sqlc.yaml                # SQLC code generator configuration
│   ├── main.go                  # Service entrypoint
│   ├── db/
│   │   ├── migrations/          # SQL schema migrations (Goose)
│   │   ├── migrations.go        # //go:embed migration loader
│   │   ├── queries/             # SQL query definitions
│   │   └── sqlc/                # Auto-generated type-safe Go code
│   └── internal/
│       ├── app/                 # Dependency injection & lifecycle wiring
│       ├── config/              # Service configuration struct
│       ├── handler/             # Gin HTTP handlers & routing
│       ├── repository/          # Database repository layer
│       └── service/             # Business logic layer
│
└── docs/                        # Specifications & Guides
    ├── PRD.md                   # Product Requirements Document
    └── WORKFLOW.md              # Guide for adding new microservices
```

---

## ❓ Troubleshooting & FAQs

### 1. "Port 5432 is already allocated"
If you already have PostgreSQL running locally:
* Open `.env.development`.
* Ensure `USER_DB_HOST_PORT=5433` (or any available host port).
* Services inside Docker will still communicate seamlessly via `db_users:5432`.

### 2. "Adding a new microservice without losing existing data"
Because each microservice has its own dedicated PostgreSQL container and isolated volume (`users_postgres_data`), you **never** need to wipe volumes (`docker compose down -v`) when adding a new service! Simply follow the [Microservice Workflow Guide](./docs/WORKFLOW.md).

### 3. "Goose version mismatch or Go toolchain errors"
Always run `make tidy` or `make build`. Dependencies are pinned and resolved within Docker to guarantee a consistent toolchain across all machines.

---

## 📖 Further Reading
* [Adding a New Microservice Guide](./docs/WORKFLOW.md)
* [Product Requirements Document (PRD)](./docs/PRD.md)
* [Architectural Guidelines & Rules](./AGENTS.md)
