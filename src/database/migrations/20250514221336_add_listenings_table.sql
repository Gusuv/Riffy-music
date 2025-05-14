-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE listenings (
    id SERIAL PRIMARY KEY,
    track_id INTEGER REFERENCES tracks(id),
    users_id INTEGER REFERENCES users(id),
    list_count INTEGER,
    list_at TIMESTAMP DEFAULT NOW()


);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
