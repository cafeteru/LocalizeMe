FROM mongo:latest

RUN apt-get update && apt-get install -y wget

COPY . /mongo-seed

RUN chmod +x /mongo-seed/start.sh
