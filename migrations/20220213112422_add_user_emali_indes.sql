-- +goose Up
CREATE UNIQUE INDEX IF NOT EXISTS users_email_uniq_idx ON "users"("email");

-- +goose Down
DROP INDEX IF EXISTS users_email_uniq_idx;

