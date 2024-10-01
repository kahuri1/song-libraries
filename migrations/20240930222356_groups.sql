-- +goose Up
CREATE TABLE IF NOT EXISTS "groups" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_group_name ON "groups" (name);

-- +goose Down
DROP INDEX IF EXISTS idx_group_name;
DROP TABLE IF EXISTS "groups";
