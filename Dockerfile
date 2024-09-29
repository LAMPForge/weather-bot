# Use the official Golang image as a build stage
FROM golang:1.20 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire application code to the container
COPY . .

# Build the Go application
RUN go build -o main .

# Use a smaller image for the final stage
FROM alpine:latest

# Set the working directory inside the final container
WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /app/main .

# Expose the port that your application uses
EXPOSE 8080

# Run the application
CMD ["./main"]
