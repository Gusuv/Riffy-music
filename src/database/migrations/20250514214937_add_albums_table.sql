-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE albums (

     id SERIAL PRIMARY KEY,
     title TEXT NOT NULL,
     artist_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
     cover_url TEXT,
     release_date TIMESTAMP DEFAULT NOW()

);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
