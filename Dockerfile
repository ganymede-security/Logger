# syntax=docker/dockerfile:1
FROM golang:1.18.3-buster

WORKDIR /home/chandlerbing98/Logger

COPY go.mod ./
COPY go.sum ./
COPY main.go ./
COPY /api ./
COPY /db ./

RUN go mod download

COPY *.go ./

RUN go build -v /home/chandlerbing/Logger/

CMD [ "/Logger" ]