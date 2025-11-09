# Stage 1: Build the Go binary
FROM golang:1.25 AS builder

WORKDIR /app

# Copy Go source
COPY . .

# Build for Linux (the container OS)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server main.go

# Stage 2: Create a minimal runtime image
FROM alpine:latest

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/server .

# Make sure it’s executable
RUN chmod +x ./server

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./server"]
