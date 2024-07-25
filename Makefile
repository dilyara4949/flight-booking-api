GOLANGCILINT ?= docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.57.2 golangci-lint
DB_URL=postgres://postgres:12345@localhost:5432/postgres?sslmode=disable
DB_URL_TEST=postgres://postgres:12345@localhost:5432/postgrestest?sslmode=disable

lint:
	  $(GOLANGCILINT) run -v --enable-all

run:
	go run cmd/main.go

create-migration:
	@read -p "migration name: " name; \
	migrate create -ext sql -dir internal/database/postgres/migration -seq $$name

migrate-up:
	migrate -database $(DB_URL) -path internal/database/postgres/migration up

migrate-down:
	migrate -database $(DB_URL) -path internal/database/postgres/migration down

migrate-up-test:
	migrate -database $(DB_URL_TEST) -path internal/database/postgres/migration up

migrate-down-test:
	migrate -database $(DB_URL_TEST) -path internal/database/postgres/migration down

testintegration:
	go test -v ./... -run TestUpdateUserHandler -tags=integration
