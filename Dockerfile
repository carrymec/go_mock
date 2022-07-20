FROM golang:latest

MAINTAINER <admin>

WORKDIR /docker

COPY . /docker

ENV GOPROXY https://goproxy.cn


RUN go build -a -o basic


ENTRYPOINT ["./basic"]
