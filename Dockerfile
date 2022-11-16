FROM golang:1.19-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR /app

ARG DB_CONNECT_STRING="host=database port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"

COPY --from=builder /app/main .

EXPOSE 80

CMD ["./main", "${DB_CONNECT_STRING}"]
