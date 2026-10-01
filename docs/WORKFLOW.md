# Microservice Integration Workflow

Follow these steps whenever you create a new microservice (e.g., `order-service`) to ensure its database and environment are correctly set up with the modern **Embedded Migrations** approach.

---

### Step 1: Define Environment Variables
Add the new service's database credentials to your `.env.development` (and `.env.example`).

```env
# Order Service Database Configuration
ORDER_DB_USER=order_dev_user
ORDER_DB_PASSWORD=your_secure_password
ORDER_DB_NAME=db_orders
```

---

### Step 2: Update Database Initialization
Add a new block to `init-db.sh` to provision the database and user for the new service.

```bash
# ... existing user-service block ...

# Order Service Setup
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "postgres" <<-EOSQL
    CREATE USER $ORDER_DB_USER WITH PASSWORD '$ORDER_DB_PASSWORD';
    CREATE DATABASE $ORDER_DB_NAME;
    GRANT ALL PRIVILEGES ON DATABASE $ORDER_DB_NAME TO $ORDER_DB_USER;
    
    \c $ORDER_DB_NAME
    ALTER SCHEMA public OWNER TO $ORDER_DB_USER;
    GRANT ALL ON SCHEMA public TO $ORDER_DB_USER;
EOSQL
```

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

### Step 4: Update `docker-compose.yml`
Add the new service to your `docker-compose.yml`. Notice that **no separate migration container is required**!

```yaml
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
    depends_on:
      db_kios:
        condition: service_healthy
```

---

### Step 5: Reset & Apply (If new database created in `init-db.sh`)
Since the database initialization script (`init-db.sh`) only runs the **first time** the database volume is initialized:

```bash
# WARNING: This deletes local docker volumes!
docker compose down -v
make up
```

---

### Step 6: Verify
1. Run `make logs` to confirm database connection and embedded migration success:
   `level=INFO msg="Embedded database migrations applied successfully"`
2. Hit the service health check endpoint (e.g. `curl http://localhost:8081/health`).
