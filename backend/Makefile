clean:
	go clean
docs:
	swagger generate spec -o swagger/swagger.json --scan-models
	swagger serve -F=swagger swagger/swagger.json --port=9090
generate-mocks:
	mockgen -source=internal/repository/UserRepository.go -destination=internal/repository/mock/UserRepositoryMock.go -package=mock
	mockgen -source=internal/core/service/UserService.go -destination=internal/core/service/mock/UserServiceMock.go -package=mock
init-project:
	make install
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install gotest.tools/gotestsum@latest
	go install github.com/go-swagger/go-swagger/cmd/swagger@latest
	go install github.com/golang/mock/mockgen@v1.6.0
install:
	go install ./...
	go mod tidy
lint:
	golangci-lint run -v
pre-commit:
	make lint
	make test
	swagger generate spec -o swagger/swagger.json --scan-models
	make uml -B
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
