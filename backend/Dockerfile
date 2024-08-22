FROM golang:latest

# Define the golangci-lint version as a build argument
ARG GOLANGCI_LINT_VERSION=v1.60.2

WORKDIR /localizeMe

COPY . /localizeMe

# Update package lists, install necessary dependencies including Git
RUN apt-get update && apt-get install -y --no-install-recommends \
    curl \
    git \
    && rm -rf /var/lib/apt/lists/*

# Install golangci-lint using the version specified by the build argument
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
    sh -s -- -b $(go env GOPATH)/bin $GOLANGCI_LINT_VERSION

# Install Go dependencies
RUN go mod download

# Run golangci-lint to check for linting issues
RUN golangci-lint run -v

# Initialize the project and generate the Swagger specification
RUN make init-project -B
RUN swagger generate spec -o swagger/swagger.json --scan-models

# Expose necessary ports
EXPOSE 8080
EXPOSE 9090

# Command to run both processes in the container using JSON format for CMD
CMD ["bash", "-c", "go run main.go & swagger serve -F=swagger swagger/swagger.json --port=9090 --no-open & wait -n"]
