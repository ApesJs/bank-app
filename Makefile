postgres:
	docker run --name postgres-app-bank -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=apesjs123 -d postgres

createdb:
	docker exec -it postgres-app-bank createdb --username=root --owner=root app-bank

dropdb:
	docker exec -it postgres-app-bank dropdb app-bank

migrateup:
	migrate -path db/migration -database "postgresql://root:apesjs123@localhost:5432/app-bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:apesjs123@localhost:5432/app-bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown