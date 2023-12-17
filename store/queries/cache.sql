-- name: GetCache :one
SELECT *
FROM cache
WHERE id = ?
LIMIT 1;

-- name: GetCaches :many
SELECT *
FROM cache;

-- name: CreateCache :one
INSERT INTO cache (id, name)
VALUES (?, ?)
RETURNING *;
