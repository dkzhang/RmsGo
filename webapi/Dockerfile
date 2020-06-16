FROM golang:latest

RUN apt-get update && apt-get install -y netcat git

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn
ENV DKZ $GOPATH/src/github.com/dkzhang

WORKDIR $DKZ
RUN git clone https://github.com/dkzhang/RmsGo.git

WORKDIR $DKZ/RmsGo
RUN go mod download

RUN chmod +x $DKZ/RmsGo/webapi/entrypoint.sh

ENTRYPOINT ["$DKZ/RmsGo/webapi/entrypoint.sh"]
