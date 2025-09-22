-- +goose Up
CREATE TABLE documents (id INTEGER PRIMARY KEY, name TEXT NOT NULL);


-- +goose Down
DROP TABLE documents;