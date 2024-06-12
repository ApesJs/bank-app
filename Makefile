postgres:
	docker run --name postgresql -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=bank_app -d postgres

createdb:
	docker exec -it postgresql createdb --username=root --owner=root bank_app

dropdb:
	docker exec -it postgresql dropdb bank_app

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank_app?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank_app?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank_app?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank_app?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 server