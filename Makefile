DB_URL=postgres://postgres:12345@localhost:5432/postgres?sslmode=disable

run:
	go run cmd/main.go

create-migration:
	@read -p "migration name: " name; \
	migrate create -ext sql -dir internal/database/postgres/migration -seq $$name

migrate-up:
	migrate -database $(DB_URL) -path internal/database/postgres/migration up

migrate-down:
	migrate -database $(DB_URL) -path internal/database/postgres/migration down
