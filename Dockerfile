# Use the official Golang image as a build stage
FROM golang:1.23.0 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum, then download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the entire application code to the container
COPY . .

# Build the Go application with static linking for compatibility with Alpine
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/main .

# Use a minimal image for the final stage
FROM alpine:latest

# Set the working directory inside the final container
WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /app/main .

# Expose the port that your application uses
EXPOSE 8080

# Run the application
CMD ["./main"]
