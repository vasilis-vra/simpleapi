# Use the official Golang image as a builder
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /simplerapi

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

COPY .env .env


# Download the dependencies
RUN go mod download

# Copy the entire application code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Use a smaller image for the final executable
FROM alpine:latest

# Set the working directory
WORKDIR /simplerapi

# Copy the binary from the builder stage
COPY --from=builder /simplerapi/main .

# Ensure the binary is executable
RUN chmod +x ./main

# Expose the port your application will run on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
