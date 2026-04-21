# Kioskecil Microservice

A modern microservice-based architecture built.

## 🛠 Tech Stack

- **Language**: Go (Golang)
- **Database**: PostgreSQL 15
- **Orchestration**: Docker & Docker Compose
- **Migrations**: Goose
- **SQL Generator**: SQLC

## 📋 Prerequisites

Ensure you have the following installed on your local machine:

- [Docker](https://www.docker.com/get-started) & Docker Compose
- [Make](https://www.gnu.org/software/make/) (optional, but recommended for convenience)

## ⚙️ Installation & Setup

1. **Clone the repository**:
   ```bash
   git clone https://github.com/PeterNex14/kioskecil-microservice.git
   cd kioskecil-microservice
   ```

2. **Configure Environment Variables**:
   Copy the example environment file and adjust values if needed:
   ```bash
   cp .env.example .env.development
   ```

3. **Start the Infrastructure**:
   Use the Makefile to build and start all containers:
   ```bash
   make up
   ```
   *Note: On the first run, this will initialize the database and run all migrations automatically.*

## 📖 Useful Commands

We use a `Makefile` to simplify common development tasks:

- `make up`: Start all services in the background.
- `make down`: Stop and remove all containers.
- `make logs`: View logs for all services.
- `make build`: Rebuild all Docker images.
- `make migrate-up`: Run pending database migrations.
- `make tidy`: Clean up Go modules in all services.

## 🏗 Project Structure

- `/user-service`: The core logic for user management.
- `/common`: Shared Go utilities and configurations used by multiple services.
- `/docs`: Detailed documentation and [workflows](./docs/WORKFLOW.md).
- `init-db.sh`: Shell script for automated database provisioning.
- `docker-compose.yml`: Infrastructure definition.

## 🤝 Contributing

For detailed instructions on how to add a new microservice to this project, please refer to the [Microservice Workflow Guide](./docs/WORKFLOW.md).
