-- +goose Up
CREATE TABLE discord_users(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username TEXT UNIQUE NOT NULL,
    discord_id INT UNIQUE NOT NULL
);


CREATE TABLE users(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    discord_id INT references discord_users(discord_id) ON DELETE CASCADE NOT NULL,
    hive_id TEXT NOT NULL,
    server TEXT NOT NULL,
    active BOOLEAN DEFAULT TRUE,
    game_uid INT NOT NULL,
    CONSTRAINT chk_server_type 
        CHECK (server IN('global', 'korea', 'japan', 'china', 'asia', 'europe')),
    UNIQUE (hive_id, server)
);


-- +goose Down
DROP TABLE users;
DROP TABLE discord_users;
