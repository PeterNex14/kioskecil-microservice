# Microservice Integration Workflow

Follow these steps whenever you create a new microservice (e.g., `order-service`) to set up its dedicated database and environment with **zero impact on existing service data**.

---

### Step 1: Define Environment Variables
Add the new service's database configuration to `.env.development` (and `.env.example`).

```env
# Order Service Database Configuration (Dedicated Container)
ORDER_DB_HOST=db_orders
ORDER_DB_PORT=5432
ORDER_DB_HOST_PORT=5434
ORDER_DB_USER=order_dev_user
ORDER_DB_PASSWORD=your_secure_password
ORDER_DB_NAME=db_orders
```

---

### Step 2: Add Dedicated Database & Service to `docker-compose.yml`
Add the dedicated PostgreSQL container and the new service to `docker-compose.yml`.

```yaml
  db_orders:
    image: postgres:15-alpine
    container_name: db_orders_container
    ports:
      - "${ORDER_DB_HOST_PORT:-5434}:5432"
    environment:
      POSTGRES_DB: ${ORDER_DB_NAME:-db_orders}
      POSTGRES_USER: ${ORDER_DB_USER:-order_dev_user}
      POSTGRES_PASSWORD: ${ORDER_DB_PASSWORD:-your_secure_password}
    volumes:
      - orders_postgres_data:/var/lib/postgresql/data
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U ${ORDER_DB_USER:-order_dev_user} -d ${ORDER_DB_NAME:-db_orders}"]
      interval: 5s
      timeout: 5s
      retries: 5

  order-service:
    build:
      context: .
      dockerfile: ./order-service/Dockerfile
    container_name: order_service_container
    ports:
      - "8081:8080"
    env_file:
      - ${ENV_FILE:-.env.development}
    environment:
      - SERVICE_NAME=order-service
      - DB_HOST=db_orders
      - DB_PORT=5432
    depends_on:
      db_orders:
        condition: service_healthy

volumes:
  users_postgres_data:
  orders_postgres_data:
```

> [!NOTE]
> Notice that **no `init-db.sh` or volume reset is needed**! Existing data in `users_postgres_data` remains completely untouched.

---

### Step 3: Embed Migrations in the New Service
In the new service's `db/` package (e.g., `order-service/db/migrations.go`):

```go
package db

import "embed"

// MigrationsFS embeds all SQL migration files inside migrations/
//
//go:embed migrations/*.sql
var MigrationsFS embed.FS
```

In the service initialization (`internal/app/app.go`):
```go
// Run Embedded Database Migrations
if cfg.AutoMigrate {
    if err := database.RunMigrations(db, order_db.MigrationsFS, "migrations"); err != nil {
        slog.Error("failed to run database migrations", "error", err)
        return nil, err
    }
}
```

---

### Step 4: Start & Verify
Run:
```bash
make up
```

1. Check logs to confirm startup and migration success:
   ```bash
   make logs
   ```
2. Verify the new service healthcheck:
   ```bash
   curl -i http://localhost:8081/health
   ```
