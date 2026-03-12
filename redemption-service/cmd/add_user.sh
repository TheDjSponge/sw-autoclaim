#!/bin/bash

NAME=$1
source ../../.env
echo "Sourced server port: ${SERVER_PORT}"
echo "{'hive_id': ${NAME}, 'server': 'china', 'discord_id': 12999321,'discord_username': ${NAME} }"

curl --location "http://localhost:${SERVER_PORT}/v1/users" \
--header 'Content-Type: application/json' \
--data "{'hive_id': '${NAME}', 'server': 'china', 'discord_id': 12999321,'discord_username': '${NAME}' }"