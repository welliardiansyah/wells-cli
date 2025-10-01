APP_NAME = wells
BIN_DIR  = bin

.PHONY: build run test clean

## Build binary wells ke folder bin/
build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) .

## Jalankan wells langsung tanpa build
run:
	go run main.go

## Jalankan unit test
test:
	go test ./...

## Bersihkan binary build
clean:
	rm -rf $(BIN_DIR)
