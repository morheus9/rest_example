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