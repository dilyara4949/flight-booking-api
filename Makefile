GOLANGCILINT ?= docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.57.2 golangci-lint
DB_URL=postgres://postgres:12345@localhost:5432/postgres?sslmode=disable
JWT_TOKEN_SECRET=my_secret_key
REST_PORT=8080
ACCESS_TOKEN_EXPIRE=877
ADDRESS=0.0.0.0
HEADER_TIMEOUT=5s
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=12345
POSTGRES_DB=postgres
POSTGRES_TIMEOUT=30
POSTGRES_MAX_CONNECTIONS=20
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=12345
REDIS_TIMEOUT=10
REDIS_TTL=5
REDIS_DATABASE=0
REDIS_POOL_SIZE=10

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
