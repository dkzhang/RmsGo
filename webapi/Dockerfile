FROM golang:latest

RUN apt-get update && apt-get install -y netcat git

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn
ENV DKZ $GOPATH/src/gitee.com/dkzhang

WORKDIR $DKZ
RUN git clone https://gitee.com/dkzhang007/RmsGo.git #20200925-0830

WORKDIR $DKZ/RmsGo
RUN go build -o rms && \
    chmod +x $DKZ/RmsGo/rms

RUN chmod +x $DKZ/RmsGo/webapi/entrypoint.sh

ENTRYPOINT ["/go/src/gitee.com/dkzhang/RmsGo/webapi/entrypoint.sh"]
