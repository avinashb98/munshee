FROM golang:alpine AS builder

RUN apk update && apk add --no-cache 'git=~2'

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

ENV PORT 8080
ENV GIN_MODE release
EXPOSE 8080

RUN env GOOS=linux GOARCH=amd64 go build -o /main

EXPOSE 8080

CMD [ "/main" ]
