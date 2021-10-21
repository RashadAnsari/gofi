export APP_NAME=gofi
export APP_VERSION=v0.0.0
export BUILD_INFO_PKG="github.com/RashadAnsari/$(APP_NAME)/pkg/version"
export LDFLAGS="-w -s -X '$(BUILD_INFO_PKG).AppVersion=$(APP_VERSION)' -X '$(BUILD_INFO_PKG).Date=$$(date)' -X '$(BUILD_INFO_PKG).BuildVersion=$$(git rev-parse HEAD | cut -c 1-8)' -X '$(BUILD_INFO_PKG).VCSRef=$$(git rev-parse --abbrev-ref HEAD)'"

all: docs format tidy lint

run-version:
	@go run -ldflags $(LDFLAGS) ./cmd/$(APP_NAME) version

run-server:
	@go run -ldflags $(LDFLAGS) ./cmd/$(APP_NAME) server

tidy:
	@go mod tidy

format:
	@find . -type f -name '*.go' -not -path './vendor/*' -exec gofmt -s -w {} +
	@find . -type f -name '*.go' -not -path './vendor/*' -exec goimports -w  -local github.com/RashadAnsari {} +

lint:
	@golangci-lint -c .golangci.yml run ./...

docs:
	@./scripts/swagger.sh

test:
	@go test -ldflags $(LDFLAGS) -v -race -p 1 ./...

ci-test:
	@go test -ldflags $(LDFLAGS) -v -race -p 1 -coverprofile=coverage.txt -covermode=atomic ./...
	@go tool cover -func coverage.txt

up:
	@docker-compose -f test/docker-compose.yml up -d

down:
	@docker-compose -f test/docker-compose.yml down
