CREATE DATABASE db_users;

CREATE USER gabs_dev WITH PASSWORD '894b08a1';

GRANT ALL PRIVILEGES ON DATABASE db_users TO gabs_dev;

\c db_users

GRANT ALL ON SCHEMA public TO gabs_dev;