#!make
# export $(shell varlock load --format env)
# POSTGRES_USER = xxx

up:
	varlock run -- docker compose up -d

createdb:
	docker compose exec postgres createdb --username=$${POSGRES_USER} simple_bank

dropdb:
	docker compose exec postgres dropdb --username=$${POSGRES_USER} simple_bank

upgrade:
	migrate -path db/migration -database ${POSTGRES_URL}?sslmode=disable -verbose up

downgrade:
	migrate -path db/migration -database ${POSTGRES_URL}?sslmode=disable -verbose down

sqlc:
	sqlc compile
	sqlc generate
xx:
	# export $(shell varlock load --format env)
	echo "wowo----${POSTGRES_USER}" && \
	echo "wowo----${POSTGRES_URL}" && \
	echo "wwwww----$$POSGRES_USER==" && \
	echo "wwwww----${APP_ENV}=="

.PHONY: up createdb dropdb upgrade downgrade xx