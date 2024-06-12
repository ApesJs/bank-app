-- name: CreateEntry :one
INSERT INTO entries (
    transfer_id,
    account_id,
    amount,
    type
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries
WHERE account_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;