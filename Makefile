#!make
include app.env


what:
	@echo "SERVER_ADDRESS=${SERVER_ADDRESS}"
	@echo "DB_SOURCE=${DB_SOURCE}"
	@echo "POSTGRES_URL=${POSTGRES_URL}"
	@echo "APP_ENV=${APP_ENV}"
up:
	docker compose up -d

createdb:
	docker compose exec postgres createdb --username=$${POSGRES_USER} simple_bank

dropdb:
	docker compose exec postgres dropdb --username=$${POSGRES_USER} simple_bank

upgrade:
	migrate -path db/migration -database ${POSTGRES_URL}?sslmode=disable -verbose up 1

downgrade:
	migrate -path db/migration -database ${POSTGRES_URL}?sslmode=disable -verbose down 1

sqlc:
	sqlc compile
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package dbmock -destination db/mock/store.go github.com/duodecanol/simplebank/db/sqlc Store


.PHONY: up createdb dropdb upgrade downgrade sqlc test server mock