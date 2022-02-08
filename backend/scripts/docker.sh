#!/bin/bash
go run main.go &
swagger serve -F=swagger swagger/swagger.json --port=9090 --no-open &
wait -n
exit $?