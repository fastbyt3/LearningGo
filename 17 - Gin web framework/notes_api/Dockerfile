# syntax=docker/dockerfile
FROM golang:latest

RUN mkdir -p /go/src/notes_api

WORKDIR /go/src/notes_api

COPY . /go/src/notes_api/

RUN go install notes_api

CMD /go/bin/notes_api

EXPOSE 8080