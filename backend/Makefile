clean:
	go clean
docs:
	swagger generate spec -o swagger/swagger.json --scan-models
	swagger serve -F=swagger swagger/swagger.json --port=9090
mocks:
	mockgen -source=internal/repository/UserRepository.go -destination=internal/repository/mock/UserRepositoryMock.go -package=mock
	mockgen -source=internal/repository/StageRepository.go -destination=internal/repository/mock/StageRepositoryMock.go -package=mock
	mockgen -source=internal/repository/LanguageRepository.go -destination=internal/repository/mock/LanguageRepositoryMock.go -package=mock
	mockgen -source=internal/repository/GroupRepository.go -destination=internal/repository/mock/GroupRepositoryMock.go -package=mock
	mockgen -source=internal/repository/BaseStringRepository.go -destination=internal/repository/mock/BaseStringRepositoryMock.go -package=mock
	mockgen -source=internal/core/service/UserService.go -destination=internal/core/service/mock/UserServiceMock.go -package=mock
	mockgen -source=internal/core/service/StageService.go -destination=internal/core/service/mock/StageServiceMock.go -package=mock
	mockgen -source=internal/core/service/LanguageService.go -destination=internal/core/service/mock/LanguageServiceMock.go -package=mock
	mockgen -source=internal/core/service/GroupService.go -destination=internal/core/service/mock/GroupServiceMock.go -package=mock
	mockgen -source=internal/core/service/BaseStringService.go -destination=internal/core/service/mock/BaseStringServiceMock.go -package=mock
init-project:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install gotest.tools/gotestsum@latest
	go install github.com/go-swagger/go-swagger/cmd/swagger@latest
	go install github.com/golang/mock/mockgen@latest
	go get github.com/jfeliu007/goplantuml/parser
	go install github.com/jfeliu007/goplantuml/cmd/goplantuml@latest
	make install
install:
	go install ./...
	go mod tidy
lint:
	golangci-lint run -v
pre-commit:
	make update
	make lint
	make mocks -B
	make test
	swagger generate spec -o swagger/swagger.json --scan-models
	make uml -B
test:
	go test `go list ./... | grep -v backend/internal/repository/mongodb/generic` -coverprofile=coverage.out
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
