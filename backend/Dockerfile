FROM golang:latest

WORKDIR /localizeMe

COPY . /localizeMe
COPY scripts/docker.sh docker.sh

RUN apt-get update
RUN make init-project -B
RUN swagger generate spec -o swagger/swagger.json --scan-models

EXPOSE 8080
EXPOSE 9090

CMD ./docker.sh
