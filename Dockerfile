# --------------------------------------
# Stage 1: Build the Go binary
# --------------------------------------
FROM golang:1.24-alpine AS builder

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Install git (nếu cần cho go mod)
RUN apk add --no-cache git

# Create working directory
WORKDIR /app

# Copy go.mod and go.sum first (cache dependencies)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary
RUN go build -o babyshop ./cmd/main.go

# Debug: list files to ensure binary exists
RUN ls -l /app

# --------------------------------------
# Stage 2: Run the Go binary
# --------------------------------------
FROM alpine:3.19

# Create /app folder explicitly
RUN mkdir -p /app
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/babyshop /app/babyshop

# Ensure executable
RUN chmod +x /app/babyshop

# Copy .env if needed
COPY .env .env

# Expose port (default Gin 8080)
EXPOSE 8080

# Run the binary
CMD ["/app/babyshop"]
