SHELL := /bin/zsh
GITHASH ?= $(shell git describe --long)
ARCH ?= $(shell uname -m)
BASE_BINARY_NAME = gsearch

## Checking OS -> naively ls /usr/local/Cellar
checkos: 
	ls -lha /usr/local/Cellar 1>/dev/null  &&
	@echo "OS OK!"

## Build gsearch binary
build: test
	@echo "Building GSEARCH binary..."
	env CGO_ENABLED=1 go build -race -ldflags "-X main.buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X main.githash=$(GITHASH)" -o ${BASE_BINARY_NAME}-${ARCH}-${GITHASH} ./cmd/
	@echo "Build done!"

## Run tests, performance tests, check for race condition
test:
	@echo "Start unit test..."
	go test -v -race ./cmd/

## Install binary
## TODO fix line 30 cp english.txt
install: checkos
	@echo "Install..."
	chmod +x ${BASE_BINARY_NAME}-* && \
	mkdir -p /usr/local/Cellar/${BASE_BINARY_NAME} && \
	cp ${BASE_BINARY_NAME}-* /usr/local/Cellar/${BASE_BINARY_NAME}/${BASE_BINARY_NAME} && \
 	cp ./configs/config.json /usr/local/Cellar/${BASE_BINARY_NAME}
	cp ./configs/english.txt /usr/local/Cellar/${BASE_BINARY_NAME}
	@echo "Set symlink..."
	ln /usr/local/Cellar/${BASE_BINARY_NAME}/${BASE_BINARY_NAME} /usr/local/bin/gsearch
	@echo "Install done!"

## uninstall:
uninstall: checkos
	@echo "Removing package... "
	@echo "Remove symlink"
	rm /usr/local/bin/gsearch
	@Recho "Remove binary"
	rm -rf /usr/local/Cellaer/${BASE_BINARY_NAME}


## build and install
all: build install


## Cleans only the binary from the project folder
clean:
	@echo "Remove old binary from the project folder"
	rm -rf gsearch-*
	@echo "Clean up system..."
	rm -rf /usr/local/bin/${BASE_BINARY_NAME}
	rm -rf /usr/local/Cellar/${BASE_BINARY_NAME}
	@echo Done!
