#!/bin/sh

echo "Running migrations..."

cd /app/migrations && goose postgres ${DB_URL} up