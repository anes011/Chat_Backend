-- +goose Up

CREATE TABLE users(
    id UUID PRIMARY KEY,
    user_name TEXT NOT NULL,
    email TEXT NOT NULL,
    pass TEXT NOT NULL,
    photo TEXT,
    created_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE users;