FROM golang:1.19 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o inventory-service ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/inventory-service .

EXPOSE 8081

CMD ["./inventory-service"]