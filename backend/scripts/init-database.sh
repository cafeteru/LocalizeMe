#!/bin/bash

# Load variables from .env file
export $(grep -v '^#' ../.env | xargs)

docker pull mongo
docker run -d --name mongo-on-docker -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=$MONGO_INITDB_ROOT_USERNAME -e MONGO_INITDB_ROOT_PASSWORD=$MONGO_INITDB_ROOT_PASSWORD mongo &
sleep 3m
docker cp ../mongo-seed/. mongo-on-docker:/
docker exec mongo-on-docker mongoimport --authenticationDatabase admin --username $MONGO_INITDB_ROOT_USERNAME --password $MONGO_INITDB_ROOT_PASSWORD --db localizeMe --collection users --type json --file users.json --jsonArray
docker exec mongo-on-docker mongoimport --authenticationDatabase admin --username $MONGO_INITDB_ROOT_USERNAME --password $MONGO_INITDB_ROOT_PASSWORD --db localizeMe --collection stages --type json --file stages.json --jsonArray
docker exec mongo-on-docker mongoimport --authenticationDatabase admin --username $MONGO_INITDB_ROOT_USERNAME --password $MONGO_INITDB_ROOT_PASSWORD --db localizeMe --collection languages --type json --file languages.json --jsonArray
docker exec mongo-on-docker mongoimport --authenticationDatabase admin --username $MONGO_INITDB_ROOT_USERNAME --password $MONGO_INITDB_ROOT_PASSWORD --db localizeMe --collection groups --type json --file groups.json --jsonArray
docker exec mongo-on-docker mongoimport --authenticationDatabase admin --username $MONGO_INITDB_ROOT_USERNAME --password $MONGO_INITDB_ROOT_PASSWORD --db localizeMe --collection baseStrings --type json --file baseStrings.json --jsonArray