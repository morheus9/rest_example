# Just easy rest app on go.
Stack
- postgress
- pgx
- tern
- slog
## Launch and testing
#### Instal dependencies:
```go mod tidy```
#### Check env file parametres in folder infra
#### Start postgres and pgadmin from docker compose or podman-compose
```
cd infra
podman-compose -f docker-compose-postgres.yaml up
```
#### For delete volumes:
```
podman-compose -f docker-compose-postgres.yaml  down --volumes
```
#### Appy migrations from tern. This is to apply migration to create the users table:
```
go install github.com/jackc/tern/v2@latest
export PATH=$PATH:$HOME/go/bin
cd migrations && tern migrate ./
```
#### Start server:
```
go run cmd/server/main.go
```
________________________________________________________________

#### For kubernetes
Create secret for postgres:
```
kubectl create secret generic db-secrets \
  --from-literal=host=postgres \
  --from-literal=port=5432 \
  --from-literal=user=myuser \
  --from-literal=password=mypassword \
  --from-literal=dbname=postgres
```
Create postgres:
```
cd infra/k8s/templates/postgres
kubectl apply -f .
```
Проверьте, что StatefulSet и Pod запущены:
```
kubectl get statefulset
kubectl get pods
```
Проверьте PersistentVolume и PersistentVolumeClaim:
```
kubectl get pv
kubectl get pvc
```
Установите приложение:
```
cd infra/k8s/templates/backend
kubectl apply -f .
```
Проверьте приложение:
```
kubectl get pods
```
#### Test the API using curl or Postman:
Можно проверить логи миграций:
```
kubectl logs myapp-95cc5b679-m2j2j -c db-migrate
```
Зайти в контейнер, например:
```
kubectl exec -it myapp-95cc5b679-66lgt -- sh
```
- Create user:
```
curl -X POST -H "Content-Type: application/json" -d '{"name": "John Doe", "email": "john@example.com"}' http://localhost:8080/users
```
- Get a user:
```
curl http://localhost:8080/users/1
```