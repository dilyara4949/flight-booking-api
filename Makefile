GOLANGCILINT ?= docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.57.2 golangci-lint
DB_URL=postgres://postgres:12345@localhost:5432/postgres?sslmode=disable

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

migrate-docker-down:
	docker-compose run app migrate -path ./internal/database/postgres/migration -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@postgres:5435/${POSTGRES_DB}?sslmode=disable" down

migrate-docker-up:
	docker-compose run app migrate -path ./internal/database/postgres/migration -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@postgres:5432/${POSTGRES_DB}?sslmode=disable" up

apply-kube:
	kubectl apply -f kubermanifests.yaml

kube-forward-port:
	kubectl port-forward services/app 8080:8080
