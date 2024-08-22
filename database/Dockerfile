FROM mongo:latest

# Copiar el script y los archivos JSON al contenedor
COPY start.sh /docker-entrypoint-initdb.d/start.sh
COPY mongo-seed/ /mongo-seed/

# Dar permisos de ejecución al script
RUN chmod +x /docker-entrypoint-initdb.d/start.sh

# Instalar dependencias adicionales (netcat, curl, dos2unix)
RUN apt-get update && apt-get install -y netcat curl dos2unix
