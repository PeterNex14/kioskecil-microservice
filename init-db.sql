CREATE USER user_dev_user WITH PASSWORD '894b08a1';

CREATE DATABASE db_users;

GRANT ALL PRIVILEGES ON DATABASE db_users TO user_dev_user;

\c db_users
ALTER SCHEMA public OWNER TO user_dev_user;
GRANT ALL ON SCHEMA public TO user_dev_user;