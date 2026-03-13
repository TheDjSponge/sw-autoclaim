#!/bin/bash

source ../../.env
echo "Sourced server port: ${SERVER_PORT}"

curl --location "http://localhost:${SERVER_PORT}/v1/coupons" \
--header 'Content-Type: application/json' \
--data '{"coupon_codes" : ["VALID_CODE","EXPIRED_CODE","INVALID_CODE","USED_CODE", "SERVER_ERROR"]}'

curl --location "http://localhost:${SERVER_PORT}/v1/users" \
--header 'Content-Type: application/json' \
--data '{"hive_id": "non_existent_user", "server": "china", "discord_id": 12366321,"discord_username": "blou" }'

curl --location "http://localhost:${SERVER_PORT}/v1/users" \
--header 'Content-Type: application/json' \
--data '{"hive_id": "random_user_28", "server": "europe", "discord_id": 12312321,"discord_username": "bla" }'

curl --location "http://localhost:${SERVER_PORT}/v1/users" \
--header 'Content-Type: application/json' \
--data '{"hive_id": "random_user_30", "server": "china", "discord_id": 22312381,"discord_username": "bli" }'