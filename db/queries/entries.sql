-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries
ORDER BY created_at
LIMIT $1
OFFSET $2;

-- name: CreateEntry :one
INSERT INTO entries (
    description, amount, category_id, type_id
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateEntry :exec
UPDATE entries
    set description = $2,
    amount = $3,
    category_id = $4,
    type_id = $5
WHERE id = $1;

-- name: DeleteEntry :exec
DELETE FROM entries
WHERE id = $1;