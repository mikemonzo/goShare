FROM golang:1.24.3-alpine3.21 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o goshare ./cmd/app

FROM alpine:3.21

WORKDIR /root/

COPY --from=builder /app/goshare .
COPY .env .

EXPOSE 8080
CMD ["./goshare"]