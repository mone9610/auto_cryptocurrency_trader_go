#!/usr/bin/env bash
#wait for the MySQL Server to come up
#sleep 15s

#run the setup script to create the DB and the schema in the DB
mysql -u root -p root auto_trader < "/docker-entrypoint-initdb.d/auto_trader.sql"
