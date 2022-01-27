#!/bin/bash

# Start the first process
go run main.go &

# Start the second process
swagger serve -F=swagger swagger/swagger.json --port=9090 --no-open &

# Wait for any process to exit
wait -n

# Exit with status of process that exited first
exit $?