FROM golang:1.19-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 80

CMD ["./main"]
