-- name: CreateSound :exec
INSERT INTO sounds (name, url)
VALUES ($1, $2);

-- name: GetSounds :many
SELECT id, name, url, created_at
FROM sounds;

-- name: GetSound :one
SELECT id, name, url, created_at
FROM sounds
WHERE id = $1;

-- name: GetRandomSound :one
SELECT id, name, url, created_at
FROM sounds
ORDER BY RANDOM()
LIMIT 1;

-- name: DeleteSound :exec
DELETE FROM sounds
WHERE id = $1;
