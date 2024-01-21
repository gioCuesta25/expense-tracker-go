-- name: GetCategory :one
SELECT * FROM categories
WHERE id = $1
ORDER BY name;

-- name: ListCategories :many
SELECT * FROM categories
ORDER BY created_at;

-- name: CreateCategory :one
INSERT INTO categories (
    name, is_for_incomes
) VALUES (
    $1, $2
)
RETURNING *;

-- name: UpdateCategory :exec
UPDATE categories
    set name = $2,
    is_for_incomes = $3
WHERE id = $1;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1;