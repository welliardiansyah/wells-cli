APP_NAME = wells
BIN_DIR  = bin

.PHONY: build run test clean

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) .

run:
	go run main.go

test:
	go test ./...

clean:
	rm -rf $(BIN_DIR)
