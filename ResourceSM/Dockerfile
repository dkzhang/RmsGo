FROM golang:latest

RUN apt-get update && apt-get install -y netcat git

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn
ENV DKZ $GOPATH/src/github.com/dkzhang

WORKDIR $DKZ
RUN git clone https://github.com/dkzhang/RmsGo.git #20200922-1016

WORKDIR $DKZ/RmsGo/ResourceSM
RUN go build -o rmsSchMet && \
    chmod +x $DKZ/RmsGo/ResourceSM/rmsSchMet

RUN chmod +x $DKZ/RmsGo/ResourceSM/entrypoint.sh

ENTRYPOINT ["/go/src/github.com/dkzhang/RmsGo/ResourceSM/entrypoint.sh"]
