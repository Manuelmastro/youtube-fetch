# Use the official Golang image to create the build artifact
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build the application
RUN go build -o main .

# Use a minimal base image for production
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose the application port (adjust if necessary)
EXPOSE 8080

# Set the entry point for the application
CMD ["./main"]
