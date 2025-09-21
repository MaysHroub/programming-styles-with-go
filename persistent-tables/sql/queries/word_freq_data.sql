-- name: GetWordsFreq :many
SELECT
    w.val AS word,
    COUNT(*) AS freq
FROM
    words w
    JOIN documents d ON d.id = w.doc_id
    LEFT JOIN stopwords sw ON sw.val = w.val
WHERE
    d.id = ?
    AND sw.val IS NULL
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


-- name: AddStopWord :exec
INSERT INTO
    stopwords (val)
VALUES
    (?);