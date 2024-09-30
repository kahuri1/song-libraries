-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE if not exists song_detail (
    id SERIAL PRIMARY KEY,
    song_id INT NOT NULL,
    text TEXT NOT NULL,
    link VARCHAR(255),
FOREIGN KEY (song_id) REFERENCES song(id) ON DELETE CASCADE
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
    drop table if exists song_detail;
-- +goose StatementEnd
