# Stage 1: Build the binary
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG SERVICE
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/service ./$SERVICE

# Stage 2: Run the binary
FROM alpine:latest

WORKDIR /root/

ARG SERVICE
COPY --from=builder /app/bin/service .

EXPOSE 8080 8081 8082 8083 8084 50051

CMD ["./service"]