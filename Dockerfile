# syntax=docker/dockerfile:1
FROM golang:1.18.3-alpine3.16 AS build-env
ADD . /logger
WORKDIR /logger
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -a -o logger


# final stage
FROM alpine:latest
WORKDIR /root/
COPY --from=build-env /logger .
EXPOSE 6000
ENTRYPOINT [ "./logger" ]