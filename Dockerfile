FROM golang:1.19-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR /app

ARG HOST
ARG PORT
ARG USER
ARG PASSWORD
ARG DB

COPY --from=builder /app/main .

EXPOSE 80

ENTRYPOINT ./main $HOST $PORT $USER $PASSWORD $DB
