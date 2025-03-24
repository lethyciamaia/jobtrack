# Usa a imagem oficial do Golang para compilar o cliente
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o client-app .

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/client-app .

# CMD ["./client-app"]