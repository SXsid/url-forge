.PHONY:  migrate run test lint
run:
	go run cmd/server/main.go

migrate:
	go run cmd/migrate/main.go

test:
	gotest ./...

lint:
	golangci-lint run

