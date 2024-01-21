CREATE TABLE IF NOT EXISTS post(
    id SERIAL PRIMARY KEY,
    title VARCHAR(500),
    text TEXT,
    post_image TEXT DEFAULT '',
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);