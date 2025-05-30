# Stage 1: Build the Go binary
FROM golang:1.24.3-alpine AS builder

# Install necessary build tools
RUN apk add --no-cache git

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Generate Swagger docs (optional, only if 'swag init' is needed inside Docker)
# RUN go install github.com/swaggo/swag/cmd/swag@latest && swag init

# Build the app statically
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# Stage 2: Minimal runtime image
FROM alpine:latest

# Install CA certificates
RUN apk add --upgrade --no-cache ca-certificates && update-ca-certificates

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/app .

# Copy Swagger docs
COPY --from=builder /app/docs ./docs

# Copy migration files
COPY --from=builder /app/migrations ./migrations

# Set environment variables
ENV PORT=8080
ENV GIN_MODE=release

# Expose the port the app runs on
EXPOSE 8080

# Run app
CMD ["./app"]
