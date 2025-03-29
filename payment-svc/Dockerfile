# Build Stage
FROM golang:1.24 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o payment-service .

# Run Stage
FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/payment-service .

EXPOSE 8082

CMD ["./payment-service"]
