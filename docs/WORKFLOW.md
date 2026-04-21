# Microservice Integration Workflow

Follow these steps whenever you create a new microservice (e.g., `order-service`) to ensure its database and environment are correctly set up.

---

### Step 1: Define Environment Variables
Add the new service's database credentials to your `.env.development` file.

```env
# Order Service Configuration
ORDER_DB_USER=order_dev_user
ORDER_DB_PASSWORD=your_secure_password
ORDER_DB_NAME=db_orders

# Goose Migration for Order Service
GOOSE_ORDER_DRIVER=postgres
GOOSE_ORDER_DBSTRING="user=order_dev_user password=your_secure_password dbname=db_orders host=db_kios port=5432 sslmode=disable"
```

---

### Step 2: Update Database Initialization
Add a new block to `init-db.sh` to create the database and user for the new service.

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

### Step 3: Update `docker-compose.yml`
Add the new service and its migration container to your `docker-compose.yml`.

1. **New Migration Service**:
   ```yaml
   migrate-order:
     build: 
       context: .
       dockerfile: order-service/Dockerfile.migration
     container_name: migrate_order_service
     env_file:
       - .env.development
     command: up
     depends_on:
       db_kios:
         condition: service_healthy
   ```

2. **New Microservice**:
   ```yaml
   order-service:
     build: ./order-service
     container_name: order_service_container
     env_file:
       - .env.development
     depends_on:
       migrate-order:
         condition: service_completed_successfully
   ```

---

### Step 4: Reset & Apply
Since the database initialization script (`init-db.sh`) only runs the **first time** the database volume is created, you must reset the volume to apply the changes to `init-db.sh`:

```bash
# WARNING: This deletes all local data in the database!
docker compose down -v
make up
```

---

### Step 5: Verify
1. Check the logs of the database to see if the new user/db were created.
2. Check the logs of the `migrate-order` service to see if migrations were applied successfully.
