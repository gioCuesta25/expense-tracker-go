-- name: CreateUser :one
INSERT INTO users (
    hashed_password, email, full_name
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: UpdateUser :exec
UPDATE users
    set hashed_password = $2,
    email = $3,
    full_name = $4
WHERE id = $1;

-- name: UpdatePassword :exec
UPDATE users
    set hashed_password = $2
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;