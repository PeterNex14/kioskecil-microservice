#!/bin/bash
set -e

# This script runs during the first container startup.
# It uses environment variables to create the application-specific database and user.

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "postgres" <<-EOSQL
    CREATE USER $USER_DB_USER WITH PASSWORD '$USER_DB_PASSWORD';
    CREATE DATABASE $USER_DB_NAME;
    GRANT ALL PRIVILEGES ON DATABASE $USER_DB_NAME TO $USER_DB_USER;
    
    \c $USER_DB_NAME
    ALTER SCHEMA public OWNER TO $USER_DB_USER;
    GRANT ALL ON SCHEMA public TO $USER_DB_USER;
EOSQL
