-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE category_relations(
    id SERIAL PRIMARY KEY,
    news_id INTEGER NOT NULL REFERENCES news(id),
    category_id INTEGER NOT NULL REFERENCES category(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
)

-- +migrate StatementEnd