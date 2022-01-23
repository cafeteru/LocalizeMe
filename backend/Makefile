clean:
	go clean
codecov:
	curl -Os https://uploader.codecov.io/latest/windows/codecov.exe
	.\codecov.exe -t f933a9e8-579a-48cb-bf96-0ab5e7621e46
docs:
	swagger generate spec -o swagger/swagger.json --scan-models
	swagger serve -F=swagger swagger/swagger.json
lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint
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
	go install gotest.tools/gotestsum@latest
	gotestsum --junitfile ../report.xml --format testname
