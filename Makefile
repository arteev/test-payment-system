VERSION=$(shell git describe --tags --match '*.*.*' --always --abbrev=0)
BUILD_TIME=$(shell date -u '+%Y-%m-%d_%I:%M:%S%p')
GIT_HASH=$(shell git rev-parse HEAD)
IMPORT_VERSION=test-payment-system/pkg/version
LDFLAGS="-w -s -X ${IMPORT_VERSION}.Version=${VERSION} -X ${IMPORT_VERSION}.DateBuild=${BUILD_TIME} -X ${IMPORT_VERSION}.GitHash=${GIT_HASH}"


DOCKER_SRC=./build/docker.local
DOCKER_COMPOSE_CMD=docker-compose -f ${DOCKER_SRC}/docker-compose.yml
DOCKER_COMPOSE_EXT_SERVICES=pg

.PHONY: build
build:
	${DOCKER_COMPOSE_CMD} build --force-rm --build-arg LDFLAGS=${LDFLAGS}

build-local:
	CGO_ENABLED=0 go build \
		-ldflags ${LDFLAGS} \
		-o ./bin/payment ./cmd/payment/main.go


run-local: stop
	${DOCKER_COMPOSE_CMD} up -d ${DOCKER_COMPOSE_EXT_SERVICES}

run: stop-soft
	${DOCKER_COMPOSE_CMD} up -d

stop:
	${DOCKER_COMPOSE_CMD} down -v --remove-orphans

stop-soft:
	${DOCKER_COMPOSE_CMD} down

generate:
	go generate ./...

.PHONE: test
test:
	go test -v `go list ./... | grep -v /internal/tests)`
	go test -v -tags=integration test-payment-system/tests
lint:
	# golangci-lint - https://github.com/golangci/golangci-lint
	#
	# golangci-lint can be installed with:
	#   curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b $(go env GOPATH)/bin v1.20.0
	#
	# More installation options: https://github.com/golangci/golangci-lint#binary-release
	#
	golangci-lint run --config .golangci.yml

# cover calculates test coverage and removes cover files
cover: cover-calc cover-clean

# cover-calc calculates test coverage
cover-calc:
	go test -coverprofile=cover.out -coverpkg=test-payment-system/... ./... && \
	grep -v -E "/proto/|/mock/|*.bindata.go" cover.out > clear_cover.out && \
	go tool cover -func=clear_cover.out

# cover-clean removes temp cover files
cover-clean:
	rm cover.out clear_cover.out


tools:
	GO111MODULE=off go get -u github.com/swaggo/swag/cmd/swag
	GO111MODULE=on go get github.com/golang/mock/mockgen@v1.4.4

swagger:
	 swag init