BUILD_DIR = bin
VERSION := $(shell git describe --tags --always --dirty)
BUILD := $(shell date +%Y-%m-%d\ %H:%M)
LDFLAGS=-ldflags="-w -s -X 'libcommon.Version=${VERSION}' -X 'libcommon.Build=${BUILD}'"

.PHONY: build

run:
	go run cmd/main.go

build:
	go build ${LDFLAGS} -o $(BUILD_DIR)/ cmd/main.go