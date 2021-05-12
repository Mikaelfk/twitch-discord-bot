# syntax=docker/dockerfile:1
FROM golang:1.16

COPY . /app
WORKDIR /app

RUN go mod download
RUN go build

CMD ["./twitch-discord-bot"]