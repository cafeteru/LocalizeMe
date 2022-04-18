# API Rest with Go (Golang)

Development of a simple Rest API with [Go](https://golang.org/).

![Go](https://img.shields.io/badge/Golang-1.18-blue.svg?logo=go&longCache=true&style=flat)

## Install project's dependencies

If you have installed `make`

```shell
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.45.2
make init-project
```

Else

```shell
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.45.2
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install gotest.tools/gotestsum@latest
go install github.com/go-swagger/go-swagger/cmd/swagger@latest
go install github.com/golang/mock/mockgen@latest
go get github.com/jfeliu007/goplantuml/parser
go get github.com/jfeliu007/goplantuml/cmd/goplantuml
go install ./...
go mod tidy
```

## Makefile

If we use Windows, we must first install 'make' with Chocolatey. (Check permissions of the chocolatey
folder `(C:\ProgramData\chocolatey`)

```shell
choco install make
```

Then to execute any step we must enter the following command:

```shell
make <command> -B # Use -B to force to execute
```

## OpenApi

We must first install Swagger dependency:

```shell
go install github.com/go-swagger/go-swagger/cmd/swagger@latest
```

Then execute the following command:

```shell
swagger generate spec -o swagger/swagger.json --scan-models
```

This action generate `swagger.json` that contains our documentation. Then we can start two types of documentation
server, Swagger or Redoc:

```shell
SwaggerUI: swagger serve -F=swagger swagger/swagger.json
Redoc: swagger serve -F=redoc swagger/swagger.json
```

These steps are configured in Makefile

### Add token to work Swagger

1. Click in this button.

![swagger_authorization_button.png](assets/img/swagger_authorization_button.png)

2. Enter the token with the text `Bearer`

![swagger_available_authorizations](assets/img/swagger_available_authorizations.png)

## How to generate a Mock from an interface

Install library:

```shell
go install github.com/golang/mock/mockgen@v1.6.0
```

```shell
mockgen -source=<source-file> -destination=<mock-file> -package=<package-name>
```

Examples:

```shell
mockgen -source=internal/repository/UserRepository.go -destination=internal/repository/mock/UserRepositoryMock.go -package=mock
mockgen -source=internal/core/service/UserService.go -destination=internal/core/service/mock/UserServiceMock.go -package=mock
```

## Create MongoDB with Docker

```shell
docker pull mongo
docker run -d --name mongo-on-docker -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=mongoadmin -e MONGO_INITDB_ROOT_PASSWORD=secret mongo &
```

To manage this database, it's recommended to install [MongoDB Compass](https://www.mongodb.com/try/download/compass)

## Docker

- Create or update image

```shell
docker build -t localize-me-backend .
```

- Execute Docker Image (create a container)

```shell
docker run --rm --name backend -p 8080:8080 -p 9090:9090 localize-me-backend &
```

- Create tag to publish image to Docker Hub

```shell
docker tag localize-me-backend igm1990/localize-me-backend:latest
```

- Publish the image to Docker Hub

```shell
docker push igm1990/localize-me-backend:latest
```

## Generate UML model

Use these commands:

```shell
go install github.com/jfeliu007/goplantuml/parser@latest
go install github.com/jfeliu007/goplantuml/cmd/goplantuml@latest
go install ./...
make uml -B
```

_Note, when you want to generate UML, you need to change ID type to `string` to remove noise created by MongoDB._

![Uml Diagram](assets/dist/UmlDiagram.svg)
