FROM postgres:13-alpine

RUN apk add --update htop

COPY ./setup.sh /docker-entrypoint-initdb.d/setup.sh

RUN chmod 0666 /docker-entrypoint-initdb.d/setup.sh