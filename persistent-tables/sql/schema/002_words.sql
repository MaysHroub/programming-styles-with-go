-- +goose Up
CREATE TABLE words (
    id INTEGER PRIMARY KEY,
    val TEXT NOT NULL,
    doc_id INT NOT NULL REFERENCES documents (id)
);


-- +goose Down
DROP TABLE words;