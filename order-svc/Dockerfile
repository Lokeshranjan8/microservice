# Build Stage
FROM golang:1.24 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o order-service .

# Run Stage
FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/order-service .

EXPOSE 8083

CMD ["./order-service"]
