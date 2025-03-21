# Build stage
FROM golang:1.23 AS builder

# Set working directory inside the container
WORKDIR /app

# Copy go mod files first for better layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code at once (using .dockerignore to exclude unwanted files)
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o backend-api ./cmd/server

# Runtime stage
FROM alpine:latest

# Add necessary packages
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/backend-api .

# Expose the port
EXPOSE 8080

# Run the application
CMD ["./backend-api"]