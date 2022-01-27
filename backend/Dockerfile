FROM golang:latest

WORKDIR /localizeMe

COPY . /localizeMe
COPY scripts/docker.sh docker.sh

RUN apt-get update
RUN make install -B
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
RUN go install gotest.tools/gotestsum@latest
RUN go install github.com/go-swagger/go-swagger/cmd/swagger@latest
RUN swagger generate spec -o swagger/swagger.json --scan-models

EXPOSE 8080
EXPOSE 9090

CMD ./docker.sh
