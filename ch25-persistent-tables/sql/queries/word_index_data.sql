-- name: AddPage :exec
INSERT INTO
    pages (number, word_id)
VALUES
    (?, ?);


-- name: GetWordPagesPairs :many
WITH
    top_words AS (
        SELECT
            w1.val AS word,
            w1.doc_id
        FROM
            words w1
        WHERE
            w1.doc_id = ?
            AND w1.val <> ''
        GROUP BY
            w1.val
        HAVING
            COUNT(*) <= 100
        ORDER BY
            w1.val ASC
        LIMIT
            ?
    )
SELECT DISTINCT
    w.val AS word,
    p.number AS page_number
FROM
    words w
    JOIN pages p ON p.word_id = w.id
    JOIN top_words t ON t.word = w.val
WHERE
    w.doc_id = t.doc_id
ORDER BY
    word ASC;