#!/bin/bash

# Definir la ruta del archivo .env
ENV_FILE="$(dirname "$0")/.env"

# Cargar las variables de entorno desde el archivo .env
if [ -f "$ENV_FILE" ]; then
  export $(grep -v '^#' "$ENV_FILE" | xargs)
else
  echo ".env file not found!"
  exit 1
fi

# URL del conector de Debezium
DEBEZIUM_URL="http://localhost:8083/connectors"

# Obtener y borrar todos los conectores existentes
existing_connectors=$(curl -s $DEBEZIUM_URL | jq -r '.[]')

if [ -n "$existing_connectors" ]; then
  echo "Deleting existing connectors..."
  for connector in $existing_connectors; do
    echo "Deleting connector: $connector"
    curl -X DELETE "$DEBEZIUM_URL/$connector"
  done
  echo "All existing connectors deleted."
else
  echo "No existing connectors found."
fi

# Definir las entidades para las que quieres crear conectores
entities=("baseStrings" "groups" "languages" "users" "permissions" "translations" "stages")

# Prefijo para los temas de Kafka
TOPIC_PREFIX="localize-me"

# Iterar sobre cada entidad y crear un conector
for entity in "${entities[@]}"; do
  echo "Creating connector for entity: $entity"

  curl -X POST -H "Content-Type: application/json" \
  --data "{
      \"name\": \"mongo-connector-$entity\",
      \"config\": {
          \"connector.class\": \"io.debezium.connector.mongodb.MongoDbConnector\",
          \"tasks.max\": \"1\",
          \"mongodb.connection.string\": \"mongodb://${MONGO_INITDB_ROOT_USERNAME}:${MONGO_INITDB_ROOT_PASSWORD}@mongo:27017\",
          \"database.include.list\": \"$MONGO_DB_NAME\",
          \"collection.include.list\": \"$MONGO_DB_NAME.$entity\",
          \"topic.prefix\": \"$TOPIC_PREFIX\",
          \"database.history.kafka.bootstrap.servers\": \"kafka:9092\",
          \"database.history.kafka.topic\": \"schema-changes.$entity\",
          \"transforms\": \"unwrap\",
          \"transforms.unwrap.type\": \"io.debezium.transforms.ExtractNewRecordState\",
          \"transforms.unwrap.drop.tombstones\": \"false\",
          \"transforms.unwrap.delete.handling.mode\": \"drop\"
      }
  }" $DEBEZIUM_URL

  echo "Connector for $entity created."
done
