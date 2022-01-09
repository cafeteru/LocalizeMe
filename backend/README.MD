# API Rest con Go (Golang)

Development of a simple Rest API with [Go](https://golang.org/).

![Go](https://img.shields.io/badge/Golang-1.17-blue.svg?logo=go&longCache=true&style=flat)

## Install project's dependencies

```shell
go get ./...
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
go get -u github.com/go-swagger/go-swagger/cmd/swagger
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

## Generate UML model

Use these commands:

```shell
make uml -B
./scripts/generateUmlDiagram.sh # Convert .puml to .svg
```

*Note, when you want to generate UML, you need to change ID type to `string` to remove noise created by MongoDB.*

![Uml Diagram](assets/dist/UmlDiagram.svg)
