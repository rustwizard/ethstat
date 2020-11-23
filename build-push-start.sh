#!/bin/bash

$(command -v docker) build -t rustwizard/ethstat-server -f build/ethstat/server/Dockerfile .

$(command -v cd) build/pg/
$(command -v docker) build -t rustwizard/postgres:13 .

$(command -v docker) login
$(command -v docker) push rustwizard/ethstat-server
$(command -v docker) push rustwizard/postgres:13

$(command -v cd) ../../
$(command -v docker-compose) -f docker-compose.yaml up