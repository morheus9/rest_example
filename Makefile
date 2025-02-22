.SILENT

build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./bin ./cmd/app/main.go

run:
	docker-compose up --remove-orphans --build server

test:
	go test ./... -coverprofile cover.out

test-coverage:
	go tool cover -func cover.out | grep total | awk '{print $3}'

create-migration:

migrate:

migrate-down: