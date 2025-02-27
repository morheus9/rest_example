# Just easy rest app on go.
Stack
- postgress
- pgx
- tern
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
- Create user:
```
curl -X POST -H "Content-Type: application/json" -d '{"name": "John Doe", "email": "john@example.com"}' http://localhost:8080/users
```
- Get a user:
```
curl http://localhost:8080/users/1
```
