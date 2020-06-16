FROM golang:latest

RUN apt-get update && apt-get install -y netcat

ENV GO111MODULE on
ENV GOPROXY GOPROXY=https://goproxy.cn

WORKDIR /go/rms

ADD go.mod .
ADD go.sum .
RUN go mod download

ADD . .

WORKDIR /rmsConfig
ADD . ./Configuration/Parameter/.

WORKDIR /go/rms
ENTRYPOINT ["/go/rms/webapi/entrypoint.sh"]
