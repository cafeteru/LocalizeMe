clean:
	go clean
docs:
	swagger generate spec -o swagger/swagger.json --scan-models
	swagger serve -F=swagger swagger/swagger.json
lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.43.0
	golangci-lint run -v
install:
	go install ./...
	go mod tidy
uml:
	goplantuml -show-aggregations internal/core/domain/ > assets/originalUmlDiagram.puml
	sed -e '/o-- "primitive.ObjectID/d' assets/originalUmlDiagram.puml > assets/UmlDiagram.puml
	rm assets/originalUmlDiagram.puml
	./scripts/generateUmlDiagram.sh
update:
	go get -u ./...
	go mod tidy
test:
	go test ./... -coverprofile=coverage.out
