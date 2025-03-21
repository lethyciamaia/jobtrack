# Multi-stage Build: 

FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum .env ./
RUN go mod download

COPY . .
RUN go build -o jobtrack-api .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/jobtrack-api .

EXPOSE ${GO_PORT}

CMD ["./jobtrack-api"]

