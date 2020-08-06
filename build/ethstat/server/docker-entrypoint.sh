#!/bin/sh
set -e

until PGPASSWORD=$ETH_PG_PASSWORD psql -h "$ETH_PG_HOST" -U "$ETH_PG_USER" -d "$ETH_PG_DB" -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing migration command"

/app/ethstat migrate up

exec "$@"