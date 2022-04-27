#! /bin/bash
docker pull mongo
docker run -d --name mongo-on-docker -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=mongoadmin -e MONGO_INITDB_ROOT_PASSWORD=secret mongo &
docker cp ../mongo-seed/. mongo-on-docker:/
docker exec mongo-on-docker mongoimport --authenticationDatabase admin --username mongoadmin --password secret --db localizeMe --collection users --type json --file users.json --jsonArray
docker exec mongo-on-docker mongoimport --authenticationDatabase admin --username mongoadmin --password secret --db localizeMe --collection stage --type json --file stages.json --jsonArray
docker exec mongo-on-docker mongoimport --authenticationDatabase admin --username mongoadmin --password secret --db localizeMe --collection languages --type json --file languages.json --jsonArray
docker exec mongo-on-docker mongoimport --authenticationDatabase admin --username mongoadmin --password secret --db localizeMe --collection groups --type json --file groups.json --jsonArray
