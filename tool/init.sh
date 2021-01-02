#!/bin/sh
docker start mysql
docker-compose exec mysql bash -c "chmod 0775 docker-entrypoint-initdb.d/init-database.sh"
docker-compose exec mysql bash -c "./docker-entrypoint-initdb.d/init-database.sh"
docker-compose exec go bash -c "go run main"