# Install dependencies
FROM docker.io/golang:1.24.1-bookworm AS deps
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Build app
FROM docker.io/golang:1.24.1-bookworm AS builder
WORKDIR /app
COPY --from=deps /go/pkg /go/pkg
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main ./cmd/server

# Install Tern
FROM docker.io/golang:1.24.1-bookworm AS tern-builder
RUN go install github.com/jackc/tern/v2@latest

# Final
FROM docker.io/debian:bookworm-slim
WORKDIR /app
RUN groupadd -r appuser && useradd -r -g appuser appuser

# Add app and tern
COPY --from=builder /app/main .
COPY --from=tern-builder /go/bin/tern /usr/local/bin/tern
# Netcat
RUN apt-get update && apt-get install -y \
    curl \
    netcat-openbsd \
    && rm -rf /var/lib/apt/lists/*

# Migrations
COPY migrations /app/migrations

RUN chown -R appuser:appuser /app
USER appuser
CMD ["./main"]