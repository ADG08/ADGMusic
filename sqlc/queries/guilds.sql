-- name: CreateGuild :exec
INSERT INTO guilds (id, name)
VALUES ($1, $2);

-- name: GetGuilds :many
SELECT id, name, created_at
FROM guilds;

-- name: GetGuild :one
SELECT id, name, created_at
FROM guilds
WHERE id = $1;