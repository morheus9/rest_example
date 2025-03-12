# Install dependencies
FROM golang:1.24.1-bookworm AS deps
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Build
FROM golang:1.24.1-bookworm AS builder
WORKDIR /app
COPY --from=deps /go/pkg /go/pkg
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main ./cmd/server

# Run
FROM debian:bookworm-slim
WORKDIR /app
RUN groupadd -r appuser && useradd -r -g appuser appuser
COPY --from=builder /app/main .
RUN chown appuser:appuser /app/main
USER appuser

CMD ["./main"]