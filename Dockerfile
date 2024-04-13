FROM golang:1.21

WORKDIR /src
COPY . .

RUN go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct &&  \
    go build -o /server

ENTRYPOINT ["/server"]
CMD ["/src/config_docker.yaml"]

