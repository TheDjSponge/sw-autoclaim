-- name: AddUser :exec
INSERT INTO Users (hive_id, server, discord_id, game_uid) 
VALUES ($1, $2, $3, $4) ON CONFLICT (hive_id, server) DO NOTHING;

-- name: DeleteUser :execresult
DELETE FROM Users 
WHERE 
    discord_id = $1 AND 
    hive_id = $2 AND  
    server = $3;

-- name: GetUserByHiveCredentials :one
SELECT * FROM Users WHERE hive_id = $1 AND server = $2;

-- name: GetAllUsers :many
SELECT * FROM Users;

