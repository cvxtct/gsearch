VERSION=0.5.0
GITHASH ?= $(shell git describe --long)

build_package:
	@echo "Building Genc binary..."
	env CGO_ENABLED=0 go build -ldflags "-X main.buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X main.githash=$(GITHASH) -X main.version=${VERSION}" -o gsearch ./cmd/
	@echo "Build done!"

install:
	cp gsearch /usr/local/bin/ && cp config_sample.json /usr/local/bin/
	@echo "Install done!"

all: build_package install