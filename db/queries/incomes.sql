-- name: GetIncomes :one
SELECT * FROM incomes
WHERE id = $1 LIMIT 1;

-- name: ListIncomes :many
SELECT * FROM incomes
ORDER BY created_at;

-- name: CreateIncomes :one
INSERT INTO incomes (
    description, amount, category_id
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: UpdateIncomes :exec
UPDATE incomes
    set description = $2,
    amount = $3,
    category_id = $4
WHERE id = $1;

-- name: DeleteIncomes :exec
DELETE FROM incomes
WHERE id = $1;