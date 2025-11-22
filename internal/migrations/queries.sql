-- name: CreateUser :one
INSERT INTO 
    users(username, email, password, created, updated)
VALUES 
    ($1, $2, $3, $4, $5) 
RETURNING *;

-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT *
FROM users
ORDER BY id;

-- name: CreateBlog :one
INSERT INTO
    blogs(title, content, user_id, created, updated)
VALUES
    ($1, $2, $3, $4, $5)
RETURNING *;

-- name: ListBlogs :many
SELECT *
FROM blogs
ORDER BY id;