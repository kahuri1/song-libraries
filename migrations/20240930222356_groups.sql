-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE if not exists "group" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table if exists "group";
-- +goose StatementEnd
