VERSION ?= $(shell git tag)
COMMIT ?=$(shell git log -1 --pretty=format:"%h")
DATE ?=$(shell git log -1 --pretty=format:"%aD")

LDFLAGS=-ldflags "-v -w -s -X 'github.com/vmedinskiy/highload/internal/cmd/server.buildVersion=${VERSION}' \
	-X 'github.com/vmedinskiy/highload/internal/cmd/server.buildDate=${DATE}' \
	-X 'github.com/vmedinskiy/highload/internal/cmd/server.buildCommit=${COMMIT}'"

BUILD        ?= go build      -o ./cmd/bin/socnet            ${LDFLAGS_SERVER} ./cmd/*.go


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