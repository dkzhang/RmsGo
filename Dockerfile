FROM golang:latest as build

ENV GO111MODULE on
ENV GOPROXY GOPROXY=https://goproxy.cn

WORKDIR /go/cache

ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/release

ADD . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o app main.go
# We can omit the symbol table, debug information and the DWARF table with the following flags: -s -w
# in order to decrease the size of executable file.

FROM scratch as prod

COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /go/release/app /
COPY --from=build /go/release/conf.yaml /

CMD ["/app"]