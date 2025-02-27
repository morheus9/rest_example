# Just easy rest app on go.
Stack
- pgx
- tern
- 
## Launch and testing
#### Instal dependencies:
```go mod tidy```
#### Check env file parametres in folder infra
#### Start postgres and pgadmin from docker compose or podman-compose
```
cd infra
podman-compose -f docker-compose-postgres.yaml  up
```
#### For delete volumes:
```
podman-compose -f docker-compose-postgres.yaml  down --volumes
```
#### Appy migrations from tern. This is to apply migration to create the users table:
```
export GOBIN="$HOME/tools"
go install github.com/jackc/tern/v2@latest
cd migratoions
tern migrate ./
```
#### Start server:
```
go run cmd/server/main.go
```
#### Test the API using curl or Postman:

###### Create user:
```
curl -X POST -H "Content-Type: application/json" -d '{"name": "John Doe", "email": "john@example.com"}' http://localhost:8080/users
```
###### Get a user:
```
curl http://localhost:8080/users/1
```



# rest_example
Rest example with stack:
- postgresql
- pgx
- tern
- slog
- swagger

#### Check .env:
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=user
DB_PASSWORD=password
DB_NAME=author_db
```
#### Start postgres + pgadmin from docker compose:
```
cd infra
docker compose up
```
#### Create a PostgreSQL database:
```
createdb 
```
#### Migrations:
```
migrate -path ./migrations -database "postgres://<user>:<pass>@localhost:5432/author_db?sslmode=disable" up
```
#### Generating from sqlc:
Define your sqlc.yaml file:
```
version: "2"
sql:
  - engine: "postgresql"
    queries: "query.sql"
    schema: "schema.sql"
    gen:
      go:
        package: "tutorial"
        out: "tutorial"
        sql_package: "pgx/v5"
```
Define your SQL schema in a schema.sql file:
```
CREATE TABLE authors (
    id BIGSERIAL PRIMARY KEY,
    name text NOT NULL,
    bio text
);
```
Define your SQL queries in a query.sql file:
```
-- name: CreateAuthor :one
INSERT INTO authors (name, bio)
VALUES ($1, $2)
RETURNING *;

-- name: GetAuthor :one
SELECT * FROM authors
WHERE id = $1;

-- name: ListAuthors :many
SELECT * FROM authors
ORDER BY id;

-- name: UpdateAuthor :one
UPDATE authors
SET name = $2, bio = $3
WHERE id = $1
RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = $1;
```
Generate:
```
sqlc generate
```
