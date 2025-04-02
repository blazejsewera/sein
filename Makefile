PROJECT="github.com/blazejsewera/sein"
SHORT_NAME="sein"
CONTAINERFILE="./Dockerfile"

PREFIX="[make]"

.PHONY: clean test

all: clean sync lint build test

build:
	@go build -o "$(SHORT_NAME)" "$(PROJECT)"
	@echo "$(PREFIX) binary $(SHORT_NAME) built"

copy-example-configs:
	@echo "$(PREFIX) additional configs for core copied"

docker-build:
	@docker build -f "$(CONTAINERFILE)" -t "$(SHORT_NAME)" .
	@echo "$(PREFIX) docker image built"

_build-binary-in-docker:
	@go build -o "/$(SHORT_NAME)" "$(PROJECT)"
	@echo "$(PREFIX) binary /$(SHORT_NAME) built for docker"

run:
	@go run "$(PROJECT)"

test:
	@go test ./...

test-race:
	@go test -race ./...

pre-commit: format
	@echo "$(PREFIX) pre-commit ok"

lint:
	@go vet ./...

format:
	@go fmt ./...

sync:
	@go mod download
	@go mod verify

tidy:
	@go mod tidy

clean:
	@rm -f "$(SHORT_NAME)"
	@go clean
	@go clean -cache
	@go clean -testcache
	@echo "$(PREFIX) binaries removed, go cache cleaned"
