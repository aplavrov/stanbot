FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o stanbot ./cmd/app

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/stanbot .

COPY config.env .

CMD ["./stanbot"]