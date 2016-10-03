BINARY=server.o
GOFILES = $(wildcard server/*.go)

.DEFAULT_GOAL := run

run: build
	./build/${BINARY}

build: clean
	@echo "****** gin mode ==> '$$GIN_MODE'"
	@echo "****** node env ==> '$$NODE_ENV'"
	@webpack
	@cp ./client/index.html ./build
	@go build -o build/${BINARY} ${GOFILES}

clean:
	@rm -rf build
