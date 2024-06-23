#!/bin/bash

MONGO_HOST=database
MONGO_AUTH_DB=admin
MONGO_USER=mongoadmin
MONGO_PASSWORD=secret

COLLECTIONS=("users" "stages" "languages" "groups" "baseStrings")
for collection in "${COLLECTIONS[@]}"; do
    mongoimport --host $MONGO_HOST --authenticationDatabase $MONGO_AUTH_DB --username $MONGO_USER --password $MONGO_PASSWORD --db localizeMe --collection $collection --type json --file /mongo-seed/$collection.json --jsonArray
done
