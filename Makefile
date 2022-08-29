start:
	go run main.go

wire:
	cd internal/app/dependency && go run github.com/google/wire/cmd/wire ./...