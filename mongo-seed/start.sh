#!/bin/bash

# Esperar a que el contenedor de la base de datos esté disponible
apt-get update
apt-get install -y netcat curl dos2unix 
curl -o wait-for-it.sh https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh 
chmod +x wait-for-it.sh 
dos2unix mongo-seed/*
./wait-for-it.sh database:27017 
./mongo-seed/import.sh
