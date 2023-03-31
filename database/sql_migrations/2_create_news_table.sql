-- +migrate Up 
-- +migrate StatementBegin

CREATE TABLE news(
    id SERIAL PRIMARY KEY, 
    user_id INTEGER NOT NULL REFERENCES users(id),
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    thumbnail TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
)

-- +migrate StatementEnd