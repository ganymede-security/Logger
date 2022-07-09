# syntax=docker/dockerfile:1
FROM golang:1.18.3-buster

ADD . /go/src/logger
WORKDIR /home/chandlerbing/logger

COPY go.mod ./
COPY go.sum ./
ADD . /home/chandlerbing/logger

RUN go mod download

COPY *.go ./

RUN go build -v .
RUN go install .

CMD [ "logger" ]