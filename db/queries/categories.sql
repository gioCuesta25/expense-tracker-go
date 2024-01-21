-- name: CreateCategory :one
INSERT INTO categories (
    name
) VALUES (
    $1
)
RETURNING *;

--name: UpdateCategory :exec
UPDATE categories
    set name = $2
WHERE id = $1;

--name: ListCategories :many
SELECT * FROM categories
ORDER BY created_at;

--name: GetCategory :one
SELECT * FROM categories
WHERE id = $1
ORDER BY name;