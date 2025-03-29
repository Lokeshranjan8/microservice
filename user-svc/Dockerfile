# Use official Go image as a base
FROM golang:1.24 AS build

# Set the working directory
WORKDIR /app

# Copy the Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary
RUN go build -o user-service .

# Use a lightweight image for production
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the compiled binary from the build stage
COPY --from=build /app/user-service .

# Expose the port (update if needed)
EXPOSE 8080

# Run the service
CMD ["./user-service"]
