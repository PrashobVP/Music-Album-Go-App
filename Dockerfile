# Step 1: Build the Go binary
FROM golang:1.21-alpine AS build

# Set the working directory inside the container

WORKDIR /app

# Copy go.mod and go.sum from the root directory to download dependencies
COPY go.mod go.sum ./

RUN go mod download

# Copy the entire source code
COPY ./cmd/api /app/cmd/api
COPY ./internal /app/internal

# Build the Go app and output the binary to the /app directory
WORKDIR /app/cmd/api
RUN go build -o /app/main .

# Step 2: Create a minimal Docker image to run the Go app
FROM alpine:3.18


# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/main /app/main

# Expose the port that your application listens on
EXPOSE 3001

# Run the binary
CMD ["/app/main"]
