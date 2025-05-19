-- name: CreateActiveChannel :exec
INSERT INTO active_channels (guild_id, channel_id)
VALUES ($1, $2);

-- name: GetActiveChannels :many
SELECT id, guild_id, channel_id, created_at
FROM active_channels;

-- name: DeleteActiveChannel :exec
DELETE FROM active_channels
WHERE id = $1;
