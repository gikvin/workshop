# Use a multi-stage build for smaller image size
FROM golang:1.22 AS builder

# Set the working directory inside the builder container
WORKDIR /app

# Copy the source code into the builder container
COPY src/* .
COPY go.mod .


# Build the Go application
RUN go mod tidy && CGO_ENABLED=0 GOOS=linux go build -o workshop .

# Use a minimal base image for the final container
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the build stage
COPY --from=builder /app/workshop .

# Set the timezone and install CA certificates
RUN apk --no-cache add ca-certificates tzdata

# Set the entrypoint command
ENTRYPOINT ["/app/workshop"]
