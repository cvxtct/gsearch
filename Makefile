VERSION=0.5.0
GITHASH ?= $(shell git describe --long)

build_package: unit_tests
	@echo "Building GSEARCH binary..."
	env CGO_ENABLED=0 go build -ldflags "-X main.buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X main.githash=$(GITHASH) -X main.version=${VERSION}" -o gsearch ./cmd/
	@echo "Build done!"

unit_tests:
	@echo "Running unit tests..."
	go test -v ./cmd/

install:
	chmod +x gsearch && cp gsearch /usr/local/bin/ && cp config.json /usr/local/bin/
	@echo "Install done!"

all: build_package install