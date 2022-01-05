clean:
	go clean
docs:
	swagger generate spec -o swagger/swagger.json --scan-models
	swagger serve -F=swagger swagger/swagger.json
test:
	go test ./... -coverprofile=coverage.out
