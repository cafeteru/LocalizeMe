clean:
	go clean
docs:
	swagger generate spec -o swagger/swagger.json --scan-models
	swagger serve -F=swagger swagger/swagger.json --port=9090
install:
	go install ./...
	go mod tidy
lint:
	golangci-lint run -v
pre-commit:
	make lint
	make test
test:
	go test ./... -coverprofile=coverage.out
	gotestsum --junitfile ../report.xml --format testname
uml:
	goplantuml -show-aggregations internal/core/domain/ > assets/originalUmlDiagram.puml
	sed -e '/o-- "primitive.ObjectID/d' assets/originalUmlDiagram.puml > assets/UmlDiagramTime.puml
	sed -e '/o-- "time.Time"/d' assets/UmlDiagramTime.puml > assets/UmlDiagram.puml
	rm assets/originalUmlDiagram.puml
	rm assets/UmlDiagramTime.puml
	./scripts/generateUmlDiagram.sh
update:
	go get -u ./...
	go mod tidy
