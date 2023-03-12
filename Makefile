VERSION ?= $(shell git tag)
COMMIT ?=$(shell git log -1 --pretty=format:"%h")
DATE ?=$(shell git log -1 --pretty=format:"%aD")

LDFLAGS=-ldflags "-v -w -s -X 'github.com/vmedinskiy/highload/internal/cmd/server.buildVersion=${VERSION}' \
	-X 'github.com/vmedinskiy/highload/internal/cmd/server.buildDate=${DATE}' \
	-X 'github.com/vmedinskiy/highload/internal/cmd/server.buildCommit=${COMMIT}'"

BUILD        ?= go build      -o ./bin/socnet            ${LDFLAGS} ./cmd/*.go


clean-generate:
	rm -rf ./api/generated
.PHONY: clean-generate

generate: clean-generate
	go generate ./...
.PHONY: generate

.PHONY: build_bin
build_bin:
	${BUILD}

.PHONY: build
build: generate build_bin

.PHONY: docker_up
docker:
	docker-compose -f build/docker-compose.yml up --build

.PHONY: docker_up_d
docker_up_d:
	docker-compose -f build/docker-compose.yml up --build -d

.PHONY: docker_down
docker_down:
	docker-compose -f build/docker-compose.yml down
	docker volume prune -f
