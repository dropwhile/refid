# environment
BUILDDIR            := ${CURDIR}/build
ARCH                := $(shell go env GOHOSTARCH)
OS                  := $(shell go env GOHOSTOS)
GOVER               := $(shell go version | awk '{print $$3}' | tr -d '.')

# app specific info
APP_VER             := $(shell git describe --always --tags|sed 's/^v//')
GITHASH             := $(shell git rev-parse --short HEAD)
GOPATH              := $(shell go env GOPATH)
VERSION_VAR         := main.ServerVersion

# flags and build configuration
GOBUILD_OPTIONS     := -trimpath
GOTEST_FLAGS        :=
GOTEST_BENCHFLAGS   :=
GOBUILD_DEPFLAGS    := -tags netgo,production
GOBUILD_LDFLAGS     ?= -s -w
GOBUILD_FLAGS       := ${GOBUILD_DEPFLAGS} ${GOBUILD_OPTIONS} -ldflags "${GOBUILD_LDFLAGS} -X ${VERSION_VAR}=${APP_VER}"

# cross compile defs
CC_BUILD_TARGETS     = reftool
CC_BUILD_ARCHES      = darwin/amd64 darwin/arm64 freebsd/amd64 linux/amd64 linux/arm64 windows/amd64
CC_OUTPUT_TPL       := ${BUILDDIR}/bin/{{.Dir}}.{{.OS}}-{{.Arch}}

# misc
DOCKER_PREBUILD     ?=

# some exported vars (pre-configure go build behavior)
export GO111MODULE=on
#export CGO_ENABLED=0
## enable go 1.21 loopvar "experiment"
export GOEXPERIMENT=loopvar
export GOOSE_DRIVER
export GOOSE_DBSTRING
export GOOSE_MIGRATION_DIR

define HELP_OUTPUT
Available targets:
  help                this help
  clean               clean up
  all                 build binaries and man pages
  check               run checks and validators
  test                run tests
  cover               run tests with cover output
  bench               run benchmarks
  build               build all binaries
endef
export HELP_OUTPUT

export PATH := "${PATH}:${GOPATH}"

.PHONY: help
help:
	@echo "$$HELP_OUTPUT"

.PHONY: clean
clean:
	@rm -rf "${BUILDDIR}"

.PHONY: setup
setup: setup-build setup-check

.PHONY: setup-build
setup-build: ${GOPATH}/bin/stringer

.PHONY: setup-check
setup-check: ${GOPATH}/bin/staticcheck ${GOPATH}/bin/gosec ${GOPATH}/bin/govulncheck

${GOPATH}/bin/staticcheck:
	go install honnef.co/go/tools/cmd/staticcheck@latest

${GOPATH}/bin/gosec:
	go install github.com/securego/gosec/v2/cmd/gosec@latest

${GOPATH}/bin/govulncheck:
	go install golang.org/x/vuln/cmd/govulncheck@latest

${GOPATH}/bin/stringer:
	go install golang.org/x/tools/cmd/stringer@latest

.PHONY: build 
build: setup-build
	@echo ">> Generating..."
	@go generate ./...
	@echo ">> Building..."
	@[ -d "${BUILDDIR}/bin" ] || mkdir -p "${BUILDDIR}/bin"
	@(for x in ${CC_BUILD_TARGETS}; do \
		echo "...$${x}..."; \
		go build ${GOBUILD_FLAGS} -o "${BUILDDIR}/bin/$${x}" ./cmd/$${x}; \
	done)
	@echo "done!"

.PHONY: test 
test:
	@echo ">> Running tests..."
	@go test -count=1 -vet=off ${GOTEST_FLAGS} ./...

.PHONY: bench
bench:
	@echo ">> Running benchmarks..."
	@go test -bench="." -run="^$$" -test.benchmem=true ${GOTEST_BENCHFLAGS} ./...

.PHONY: cover
cover:
	@echo ">> Running tests with coverage..."
	@go test -vet=off -cover ${GOTEST_FLAGS} ./...

.PHONY: check
check: setup-check
	@echo ">> Running checks and validators..."
	@echo "... staticcheck ..."
	@${GOPATH}/bin/staticcheck ./...
	@echo "... go-vet ..."
	@go vet ./...
	@echo "... gosec ..."
	@${GOPATH}/bin/gosec -quiet ./...
	@echo "... govulncheck ..."
	@${GOPATH}/bin/govulncheck ./...

.PHONY: update-go-deps
update-go-deps:
	@echo ">> updating Go dependencies..."
	@go get -u all
	@go mod tidy

.PHONY: all
all: build
