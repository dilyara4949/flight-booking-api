GOLANGCILINT ?= docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.57.2 golangci-lint
DB_URL=postgres://postgres:12345@localhost:5432/postgres?sslmode=disable
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=12345
REDIS_TIMEOUT=10
REDIS_LONG_CACHE_DURATION=5m
REDIS_SHORT_CACHE_DURATION=5m
REDIS_DATABASE=0
REDIS_POOL_SIZE=10

lint:
	  $(GOLANGCILINT) run -v

run:
	go run cmd/main.go

create-migration:
	@read -p "migration name: " name; \
	migrate create -ext sql -dir internal/database/postgres/migration -seq $$name

migrate-up:
	migrate -database $(DB_URL) -path internal/database/postgres/migration up

migrate-down:
	migrate -database $(DB_URL) -path internal/database/postgres/migration down

export_env:
	@echo "export REDIS_HOST=$(REDIS_HOST)" >> set_env.sh
	@echo "export REDIS_PORT=$(REDIS_PORT)" >> set_env.sh
	@echo "export REDIS_PASSWORD=$(REDIS_PASSWORD)" >> set_env.sh
	@echo "export REDIS_TIMEOUT=$(REDIS_TIMEOUT)" >> set_env.sh
	@echo "export REDIS_SHORT_CACHE_DURATION=$(REDIS_SHORT_CACHE_DURATION)" >> set_env.sh
	@echo "export REDIS_LONG_CACHE_DURATION=$(REDIS_LONG_CACHE_DURATION)" >> set_env.sh
	@echo "export REDIS_DATABASE=$(REDIS_DATABASE)" >> set_env.sh
	@echo "export REDIS_POOL_SIZE=$(REDIS_POOL_SIZE)" >> set_env.sh
	@chmod +x set_env.sh