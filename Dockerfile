# syntax=docker/dockerfile:1
FROM golang:1.18.3-buster

ADD . /go/src/logger
WORKDIR /home/chandlerbing/Logger

COPY go.mod ./
COPY go.sum ./
ADD . /home/chandlerbing/Logger

RUN go mod download

COPY *.go ./

RUN go build -v .
RUN go install .

CMD [ "logger" ]