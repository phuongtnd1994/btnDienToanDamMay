#!/bin/sh

set -e

# echo "run db migration"
# /app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up
source /app/app.env

echo "start the app"
exec "$@" #Take all parameters passed to the script and execute them as a command