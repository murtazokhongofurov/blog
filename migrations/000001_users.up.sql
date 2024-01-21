CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(255),
    login VARCHAR(255),
    password TEXT,
    photo TEXT DEFAULT '',
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);