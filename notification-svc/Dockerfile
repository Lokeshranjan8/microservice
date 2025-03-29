FROM golang:1.21 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o notification-service .

FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/notification-service .

EXPOSE 8081

CMD ["./notification-service"]
