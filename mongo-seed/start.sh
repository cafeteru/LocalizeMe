#!/bin/bash

WAIT_TIME=30 # seconds

echo "Waiting $WAIT_TIME seconds for MongoDB to become available..."

sleep $WAIT_TIME

echo "Wait time completed. Starting data import..."

MONGO_HOST="mongo"
MONGO_AUTH_DB="admin"
MONGO_USER="mongoadmin"
MONGO_PASSWORD="secret"

COLLECTIONS=("users" "stages" "languages" "groups" "baseStrings")
for collection in "${COLLECTIONS[@]}"; do
    mongoimport --host $MONGO_HOST --authenticationDatabase $MONGO_AUTH_DB --username $MONGO_USER --password $MONGO_PASSWORD --db localizeMe --collection $collection --type json --file ./mongo-seed/data/$collection.json --jsonArray
done

echo "Data import completed."
