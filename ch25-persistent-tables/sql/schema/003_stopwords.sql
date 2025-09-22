-- +goose Up
CREATE TABLE stopwords (
    id INTEGER PRIMARY KEY, 
    val TEXT NOT NULL
);


-- +goose Down
DROP TABLE stopwords;