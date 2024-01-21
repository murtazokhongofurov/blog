CREATE TABLE IF NOT EXISTS comment(
    id SERIAL PRIMARY KEY,
    post_id INT NOT NULL REFERENCES post(id),
    user_id INT NOT NULL REFERENCES users(id),
    text TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);