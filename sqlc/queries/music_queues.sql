-- name: CreateMusicQueue :exec
INSERT INTO music_queues (guild_id, channel_id, url, title, added_by)
VALUES ($1, $2, $3, $4, $5);

-- name: GetMusicQueues :many
SELECT id, guild_id, channel_id, url, title, added_by, created_at
FROM music_queues;

-- name: GetMusicQueue :one
SELECT id, guild_id, channel_id, url, title, added_by, created_at
FROM music_queues
WHERE id = $1;