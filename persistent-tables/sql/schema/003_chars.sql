-- +goose Up
CREATE TABLE chars (
    id INTEGER PRIMARY KEY,
    val TEXT NOT NULL,
    word_id INT NOT NULL REFERENCES words (id)
);


-- +goose Down
DROP TABLE chars;