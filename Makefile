sqlc:
	docker pull kjconroy/sqlc

sqlcinit:
	docker run --rm -v E:\MonkeyCode\github.com\speauty\backend:/src -w /src kjconroy/sqlc init

sqlcgenerate:
	docker run --rm -v E:\MonkeyCode\github.com\speauty\backend:/src -w /src kjconroy/sqlc generate

postgres:
	docker run -d -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root --name db_postgre_v13 postgres:13-alpine

createdb:
	docker exec -it db_postgre_v13 createdb --username=root --owner=root backend

dropdb:
	docker exec -it db_postgre_v13 dropdb backend

test:
	go test -v -cover ./...

mockStore:
	mockgen -package mockdb -destination src/db/mock/store.go github.com/speauty/backend/src/db/sqlc Store

migratecreate:
	migrate create -ext sql -dir src/db/migration -seq add_users

migrateup:
    migrate -path src/db/migration -database "postgresql://root:root@localhost:5433/backend?sslmode=disable" -verbose up

migratedown:
    migrate -path src/db/migration -database "postgresql://root:root@localhost:5433/backend?sslmode=disable" -verbose down
