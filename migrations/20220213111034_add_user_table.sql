-- +goose Up
CREATE TABLE IF NOT EXISTS users
(
    id         SERIAL PRIMARY KEY,
    email      varchar(255) NOT NULL,
    password   varchar(255) NOT NULL,
    status     int default (0),
    created_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE
);

-- +goose Down
DROP TABLE IF EXISTS users;

