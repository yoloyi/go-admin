.PHONY: start, wire
name = admin

start:
	go run ./cmd/main.go --c ./cmd/config.yml

wire:
	wire gen ./internal/models/repositories
	wire gen ./internal/services


build:
	go build -o $(name) -ldflags="-s -w" ./cmd/main.go