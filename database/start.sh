#!/bin/bash

apt-get update
apt-get install -y netcat curl dos2unix

HOST="database"
PORT="27017"
TIMEOUT=15

if nc -z "$HOST" "$PORT"; then
    echo "Host $HOST:$PORT is available after $TIMEOUT seconds"
else
    echo "Host $HOST:$PORT is not available within $TIMEOUT seconds"
fi

MONGO_HOST=$HOST
MONGO_AUTH_DB=admin
MONGO_USER=mongoadmin
MONGO_PASSWORD=secret

COLLECTIONS=("users" "stages" "languages" "groups" "baseStrings")
for collection in "${COLLECTIONS[@]}"; do
    mongoimport --host $MONGO_HOST --authenticationDatabase $MONGO_AUTH_DB --username $MONGO_USER --password $MONGO_PASSWORD --db localizeMe --collection $collection --type json --file /mongo-seed/$collection.json --jsonArray
done