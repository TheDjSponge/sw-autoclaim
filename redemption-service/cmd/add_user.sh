#!/bin/bash

NAME=$1
source ../../.env
echo "Sourced server port: ${SERVER_PORT}"
echo 'Added user: {"hive_id": "'$1'", "server": "china", "discord_id": 22378381,"discord_username": "'$1'" }'

curl --location "http://localhost:${SERVER_PORT}/v1/users" \
--header 'Content-Type: application/json' \
--data '{"hive_id": "'$1'", "server": "china", "discord_id": 22378381,"discord_username": "'$1'" }'