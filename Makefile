GITHASH ?= $(shell git describe --long)
ARCH ?= $(shell uname -m)

## build gsearch binary
build: test
	@echo "Building GSEARCH binary..."
	env CGO_ENABLED=1 go build -race -ldflags "-X main.buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X main.githash=$(GITHASH)" -o gsearch-${ARCH}-${GITHASH} ./search-service/cmd/api
	@echo "Build done!"

## run tests, performance tests, check for race condition
test:
	@echo "Start unit test..."
	go test -v -race ./search-service/cmd/api/

## install program
install:
	@echo "Install..."
	chmod +x gsearch && cp gsearch /usr/local/bin/ && cp ./config/config.json /usr/local/bin/
	@echo "Install done!"

## build and install
all: build install
