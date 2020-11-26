.PHONY: start

start:
	go run ./cmd/main.go --c ./cmd/config.yml

wire:
	wire gen ./internal/models/repositories
	wire gen ./internal/services