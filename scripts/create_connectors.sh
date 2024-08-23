#!/bin/bash

# Cargar las variables de entorno desde el archivo .env
if [ -f .env ]; then
  export $(cat .env | xargs)
else
  echo ".env file not found!"
  exit 1
fi

# Definir las entidades para las que quieres crear conectores
entities=("BaseString" "Group" "Language" "User" "Permission" "Translation" "Stage")

# URL del conector de Debezium
DEBEZIUM_URL="http://localhost:8083/connectors"

# Iterar sobre cada entidad y crear un conector
for entity in "${entities[@]}"; do
  echo "Creating connector for entity: $entity"

  curl -X POST -H "Content-Type: application/json" \
  --data '{
      "name": "mongo-connector-'$entity'",
      "config": {
          "connector.class": "io.debezium.connector.mongodb.MongoDbConnector",
          "tasks.max": "1",
          "mongodb.hosts": "mongo:27017",
          "mongodb.name": "'$MONGO_DB_NAME'",
          "mongodb.user": "'$MONGO_INITDB_ROOT_USERNAME'",
          "mongodb.password": "'$MONGO_INITDB_ROOT_PASSWORD'",
          "collection.include.list": "'$MONGO_DB_NAME'.'$entity'",
          "database.history.kafka.bootstrap.servers": "kafka:9092",
          "database.history.kafka.topic": "schema-changes.'$entity'",
          "transforms": "unwrap",
          "transforms.unwrap.type": "io.debezium.transforms.ExtractNewRecordState",
          "transforms.unwrap.drop.tombstones": "false",
          "transforms.unwrap.delete.handling.mode": "drop"
      }
  }' $DEBEZIUM_URL

  echo "Connector for $entity created."
done
