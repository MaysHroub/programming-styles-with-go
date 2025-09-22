-- name: GetWordsFreq :many
SELECT
    w.val AS word,
    COUNT(*) AS freq
FROM
    words w
    LEFT JOIN stopwords sw ON sw.val = w.val
WHERE
    w.doc_id = ?
    AND sw.val IS NULL
GROUP BY
    w.val
ORDER BY
    freq DESC
LIMIT
    ?;


-- a. 25 most frequent words per book
-- name: GetWordsFreqPerDoc :many
SELECT
    w.doc_id,
    w.val AS word,
    COUNT(*) AS freq
FROM
    words w
    LEFT JOIN stopwords sw ON sw.val = w.val
WHERE
    sw.val IS NULL
GROUP BY
    w.doc_id,
    w.val
ORDER BY
    freq DESC
LIMIT
    ?;


-- b. Word count per book  
-- name: GetWordsCountPerDoc :many
SELECT
    doc_id,
    COUNT(*) AS words_count
FROM
    words
GROUP BY
    doc_id;


-- c. Character count per book
-- name: GetCharsCountPerDoc :many
SELECT
    doc_id,
    SUM(LENGTH(val)) AS chars_count
FROM
    words
GROUP BY
    doc_id;


-- d. Longest word per book
-- name: GetLongestWordsPerDoc :many
WITH
    ranked_words AS (
        SELECT
            doc_id,
            val,
            RANK() OVER (
                PARTITION BY
                    doc_id
                ORDER BY
                    LENGTH(val) DESC
            ) AS rnk
        FROM
            words
    )
SELECT
    doc_id,
    val AS longest_word
FROM
    ranked_words
WHERE
    rnk = 1;


-- e. Average word length
-- name: GetAvgWordLength :one
SELECT
    AVG(LENGTH(val))
FROM
    words;


-- f. Combined length of characters in the top 25 words of each book (this one gemini made it)
-- name: GetCombinedLengthOfTop25WordsPerDoc :many
WITH
    word_frequencies AS (
        -- First, count the frequency of each word in each document
        SELECT
            w.doc_id,
            w.val AS word,
            COUNT(*) AS freq
        FROM
            words w
            LEFT JOIN stopwords sw ON sw.val = w.val
        WHERE
            sw.val IS NULL
        GROUP BY
            w.doc_id,
            w.val
    ),
    ranked_words AS (
        -- Then, rank words by frequency within each document
        SELECT
            doc_id,
            word,
            ROW_NUMBER() OVER (
                PARTITION BY
                    doc_id
                ORDER BY
                    freq DESC
            ) AS rn
        FROM
            word_frequencies
    )
    -- Finally, filter for the top 25 words in each document and calculate the sum of their lengths
SELECT
    doc_id,
    SUM(LENGTH(word)) AS combined_length
FROM
    ranked_words
WHERE
    rn <= 25
GROUP BY
    doc_id;


-- name: GetAllDocIDs :many
SELECT
    id
FROM
    documents;


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