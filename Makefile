BINARY_NAME=task-logger.exe
DIST_DIR=dist

.PHONY: all build clean test run

all: build

build:
	go build -o $(DIST_DIR)/$(BINARY_NAME) main.go

clean:
	@[ -n "$(DIST_DIR)" ] && $(RM) -r "$(DIST_DIR)" || true

test:
	go test ./...

run:
	go run main.go
