FROM mongo:latest

# Copy the script and JSON files into the container
COPY start.sh /docker-entrypoint-initdb.d/start.sh
COPY mongo-seed/ /mongo-seed/

# Grant execution permissions to the script
RUN chmod +x /docker-entrypoint-initdb.d/start.sh

# Install additional dependencies (netcat, curl, dos2unix)
RUN apt-get update && apt-get install -y netcat curl dos2unix
