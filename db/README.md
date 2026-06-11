### Run from the dir ./db to migrate the database

`migrate -path ./migrations -database "postgres://admin:admin@localhost:5432/exchange?sslmode=disable" up`

### Run to generate the sqlc functions

`sqlc generate`
