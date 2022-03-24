# syntax=docker/dockerfile:1
FROM golang:1.16-alpine


WORKDIR /app

COPY *.go ./

RUN go mod init admin
RUN go mod tidy


RUN go build -o /docker-admin

CMD [ "/docker-admin" ]
