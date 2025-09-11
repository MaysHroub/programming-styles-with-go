-- name: GetWordsFreq :many
SELECT
    w.val AS word,
    COUNT(*) AS freq
FROM
    words w
GROUP BY
    w.val
ORDER BY
    freq DESC;


-- name: AddDocument :one
INSERT INTO
    documents (name)
VALUES
    (?)
RETURNING
    *;


-- name: AddWord :one
INSERT INTO
    words (val, doc_id)
VALUES
    (?, ?)
RETURNING
    *;


-- name: AddChar :one
INSERT INTO
    chars (val, word_id)
VALUES
    (?, ?)
RETURNING
    *;
