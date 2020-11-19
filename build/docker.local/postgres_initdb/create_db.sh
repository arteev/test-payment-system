#!/bin/bash

#
# The script uses POSTGRES_MULTIPLE_DATABASES environment variable to create
# databases provided in the variable as comma separated list. See
# 'docker-compose.yml' file for example.
#
# By default 'pg' docker image can creates one database and this script used
# to create database inside a container it's placed in
# '/docker-entrypoint-initdb.d'
#

set -e

function create_user_and_database() {
    local database=$1
    echo "  Creating database '$database'"
    psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
	    CREATE DATABASE $database;
	    GRANT ALL PRIVILEGES ON DATABASE $database TO $POSTGRES_USER;
EOSQL
}

if [ -n "$POSTGRES_MULTIPLE_DATABASES" ]; then
    echo "Multiple database creation requested: $POSTGRES_MULTIPLE_DATABASES"
    for db in $(echo $POSTGRES_MULTIPLE_DATABASES | tr ',' ' '); do
        create_user_and_database $db
    done
    echo "Multiple databases created"
fi
