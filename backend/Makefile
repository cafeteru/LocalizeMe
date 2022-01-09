clean:
	go clean
docs:
	swagger generate spec -o swagger/swagger.json --scan-models
	swagger serve -F=swagger swagger/swagger.json
install:
	go get ./...
	go mod tidy
uml:
	 goplantuml -show-aggregations internal/core/domain/ > assets/UmlDiagram.puml
test:
	go test ./... -coverprofile=coverage.out
