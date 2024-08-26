all: build run

build:
	go build -o ./bin/huffman-go ./cmd/main.go
run:
	./bin/huffman-go
