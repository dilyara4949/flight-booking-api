GOLANGCILINT ?= docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.57.2 golangci-lint
DB_URL=postgres://postgres:12345@localhost:5432/postgres?sslmode=disable

lint:
	  $(GOLANGCILINT) run -v --enable-all

run:
	go run cmd/main.go

create-migration:
	@read -p "migration name: " name; \
	timestamp=$$(date +%Y%m%d%H%M); \
	touch internal/database/postgres/migration/$${timestamp}_$$name.up.sql; \
	touch internal/database/postgres/migration/$${timestamp}_$$name.down.sql

migrate-up:
	migrate -database $(DB_URL) -path internal/database/postgres/migration up

migrate-down:
	migrate -database $(DB_URL) -path internal/database/postgres/migration down
