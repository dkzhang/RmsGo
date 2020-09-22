FROM golang:latest

RUN apt-get update && apt-get install -y netcat git

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn
ENV DKZ $GOPATH/src/gitee.com/dkzhang

WORKDIR $DKZ
RUN git clone https://gitee.com/dkzhang007/RmsGo.git #20200922-1406

WORKDIR $DKZ/RmsGo/ResourceSM
RUN go build -o schmet && \
    chmod +x $DKZ/RmsGo/ResourceSM/schmet

RUN chmod +x $DKZ/RmsGo/ResourceSM/entrypoint.sh

ENTRYPOINT ["/go/src/gitee.com/dkzhang/RmsGo/ResourceSM/entrypoint.sh"]
