#!/bin/sh
set -e

until PGPASSWORD=$PG_PASSWORD psql -h "$PG_HOST" -U "$PG_USER" -d "$PG_NAME" -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing migration command"

/app/ethstat migrate up

exec "$@"