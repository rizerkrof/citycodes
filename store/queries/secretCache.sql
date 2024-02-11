-- name: GetSecretCache :one
SELECT * FROM secret_cache WHERE id = ? LIMIT 1;

-- name: GetSecretCaches :many
SELECT * FROM secret_cache;

-- name: CreateSecretCache :one
INSERT INTO
    secret_cache (id, name, image_url)
VALUES (?, ?, ?) RETURNING *;

-- name: PatchSecretCacheImageUrl :one
UPDATE secret_cache SET image_url = ? WHERE id = ? RETURNING *;