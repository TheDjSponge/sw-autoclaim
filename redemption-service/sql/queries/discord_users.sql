-- name: AddDiscordUser :exec
INSERT INTO discord_users (username, discord_id) 
VALUES ($1, $2) ON CONFLICT (discord_id) DO NOTHING;

