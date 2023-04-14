# Builder stage
FROM golang:1.20.3-alpine3.17 AS builder

WORKDIR /app
ADD . .
RUN go mod download
RUN go build -o main main.go

# Development stage
FROM alpine:3.17 AS dev

WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080

CMD [ "/app/main", "migrate" ]
