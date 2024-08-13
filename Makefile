GOLANGCILINT ?= docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.57.2 golangci-lint
DB_URL=postgres://postgres:12345@localhost:5432/postgres?sslmode=disable
DB_URL_TEST=postgres://postgres:12345@localhost:5432/postgrestest?sslmode=disable
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=12345
POSTGRES_DB_TEST=postgrestest
POSTGRES_TIMEOUT=30
POSTGRES_MAX_CONNECTIONS=20
JWT_TOKEN_SECRET=my_secret_key
REST_PORT=8080
ACCESS_TOKEN_EXPIRE=877
ADDRESS=0.0.0.0

.PHONY: export_env

lint:
	  $(GOLANGCILINT) run -v

run:
	go run cmd/main.go

create-migration:
	@read -p "migration name: " name; \
	migrate create -ext sql -dir internal/database/postgres/migration -tz "UTC" $$name

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

export_env:
	@echo "export POSTGRES_HOST=$(POSTGRES_HOST)" > set_env.sh
	@echo "export POSTGRES_PORT=$(POSTGRES_PORT)" >> set_env.sh
	@echo "export POSTGRES_USER=$(POSTGRES_USER)" >> set_env.sh
	@echo "export POSTGRES_PASSWORD=$(POSTGRES_PASSWORD)" >> set_env.sh
	@echo "export POSTGRES_DB_TEST=$(POSTGRES_DB_TEST)" >> set_env.sh
	@echo "export POSTGRES_TIMEOUT=$(POSTGRES_TIMEOUT)" >> set_env.sh
	@echo "export POSTGRES_MAX_CONNECTIONS=$(POSTGRES_MAX_CONNECTIONS)" >> set_env.sh
	@echo "export JWT_TOKEN_SECRET=$(JWT_TOKEN_SECRET)" >> set_env.sh
	@echo "export REST_PORT=$(REST_PORT)" >> set_env.sh
	@echo "export ACCESS_TOKEN_EXPIRE=$(ACCESS_TOKEN_EXPIRE)" >> set_env.sh
	@echo "export ADDRESS=$(ADDRESS)" >> set_env.sh
	@chmod +x set_env.sh