-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE comment(
    id SERIAL PRIMARY KEY, 
    news_id INTEGER NOT NULL REFERENCES news(id),
    user_id INTEGER NOT NULL REFERENCES users(id),
    comment TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
)

-- +migrate StatementEnd