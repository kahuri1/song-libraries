-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS songs (
    id SERIAL PRIMARY KEY,
    group_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    release_date DATE NOT NULL,
FOREIGN KEY (group_id) REFERENCES "groups"(id) ON DELETE CASCADE
    );

CREATE INDEX idx_songs_title ON songs(title);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';.
    drop table if exists songs;
-- +goose StatementEnd
