FROM golang:latest

RUN apt-get update && apt-get install -y netcat git

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn

WORKDIR /go/src
RUN git clone https://github.com/dkzhang/RmsGo.git

WORKDIR /go/src/RmsGo
RUN go mod download

ENTRYPOINT ["/go/src/RmsGo/webapi/entrypoint.sh"]
