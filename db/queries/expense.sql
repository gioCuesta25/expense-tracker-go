-- name: GetExpense :one
SELECT * FROM expenses
WHERE id = $1 LIMIT 1;

-- name: ListExpenses :many
SELECT * FROM expenses
ORDER BY created_at;

-- name: CreateExpense :one
INSERT INTO expenses (
    description, amount, category_id
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: UpdateExpense :exec
UPDATE expenses
    set description = $2,
    amount = $3,
    category_id = $4
WHERE id = $1;

-- name: DeleteExpense :exec
DELETE FROM expenses
WHERE id = $1;