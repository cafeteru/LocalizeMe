#!/bin/bash

# Directorio donde se guardarán los logs
LOG_DIR="./logs"

# Crear el directorio si no existe
mkdir -p $LOG_DIR

# Obtener la lista de nombres de todos los contenedores, incluyendo los que están en estado exited
containers=$(docker ps -a --format '{{.Names}}')

# Recorrer cada contenedor y guardar sus logs en un archivo
for container in $containers
do
    log_file="${LOG_DIR}/${container}.log"
    docker logs $container &> $log_file
done
