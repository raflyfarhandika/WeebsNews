-- +migrate Up 
-- +migrate StatementBegin
CREATE TYPE roles AS ENUM ('admin', 'user');

CREATE TABLE users(
    id SERIAL PRIMARY KEY, 
    first_name VARCHAR(255),
    last_name VARCHAR (255),
    username VARCHAR(32) UNIQUE NOT NULL, 
    password VARCHAR(255) NOT NULL,
    email VARCHAR(256) NOT NULL UNIQUE,
    role roles NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(), 
    updated_at TIMESTAMP DEFAULT NOW()
)

-- +migrate StatementEnd