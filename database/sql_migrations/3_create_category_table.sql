-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE category(
    id SERIAL PRIMARY KEY,
    category_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
)

-- +migrate StatementEnd