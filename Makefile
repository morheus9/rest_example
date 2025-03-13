build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./bin ./cmd/server/main.go

run:
#cd infra && docker-compose -f docker-compose-postgres.yaml up --remove-orphans 
	cd infra && podman-compose -f docker-compose-postgres.yaml up --remove-orphans

test:
	go test ./... -coverprofile cover.out

test-coverage:
	go tool cover -func cover.out | grep total | awk '{print $3}'

create-migration:

migrate:
	go install github.com/jackc/tern/v2@latest
	export PATH=$PATH:$HOME/go/bin
	cd migrations && tern migrate ./

migrate-down:
	cd migrations && tern migrate --destination 0 ./