FROM golang:latest

RUN apt-get update && apt-get install -y netcat git

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn
ENV DKZ $GOPATH/src/github.com/dkzhang

WORKDIR $DKZ
RUN git clone https://github.com/dkzhang/RmsGo.git #20200826-0925

WORKDIR $DKZ/RmsGo
RUN go build -o rms && \
    chmod +x $DKZ/RmsGo/rms

RUN chmod +x $DKZ/RmsGo/webapi/entrypoint.sh

ENTRYPOINT ["/go/src/github.com/dkzhang/RmsGo/webapi/entrypoint.sh"]
