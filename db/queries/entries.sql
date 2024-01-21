-- name: GetIncome :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: ListIncomes :many
SELECT * FROM entries
ORDER BY created_at;

-- name: CreateIncome :one
INSERT INTO entries (
    description, amount, category_id, type_id
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateIncome :exec
UPDATE entries
    set description = $2,
    amount = $3,
    category_id = $4,
    type_id = $5
WHERE id = $1;

-- name: DeleteIncome :exec
DELETE FROM entries
WHERE id = $1;