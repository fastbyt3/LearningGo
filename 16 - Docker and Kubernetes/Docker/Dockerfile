# syntax=docker/dockerfile:1
FROM golang:latest

RUN mkdir /app
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /wapp

EXPOSE 8080

ENTRYPOINT ["/wapp"]
