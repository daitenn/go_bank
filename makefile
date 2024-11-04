postgres:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17-alpine

create-db:
	docker exec -it postgres17 createdb --username=root --owner=root simple_bank

drop-db:
	docker exec -it postgres17 createdb --username=root --owner=root simple_bank

migrate-up:
	migrate -path db/migration/ -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrate-down:
	migrate -path db/migration/ -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

fmt:
	go fmt ./...

.PHONY: postgres create-db drop-db migrate-up migrate-down sqlc test