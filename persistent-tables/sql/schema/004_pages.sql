-- +goose Up
CREATE TABLE pages (
    id INTEGER PRIMARY KEY,
    number INTEGER NOT NULL,
    word_id INTEGER NOT NULL REFERENCES words (id)
);


-- +goose Down
DROP TABLE pages;