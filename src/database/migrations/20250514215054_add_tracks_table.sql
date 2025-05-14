-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE tracks (

    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    artist_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    file_url TEXT NOT NULL,
    cover_url TEXT,
    duration INTEGER NOT NULL,
    publication_date TIMESTAMP DEFAULT NOW(),
    genre_id INTEGER REFERENCES genres(id),
    album_id INTEGER REFERENCES albums(id)

);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
