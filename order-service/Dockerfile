FROM golang:1.19 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download 

COPY . .

RUN go build -o order-service ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/order-service .

EXPOSE 8082

CMD [ "./order-service" ]