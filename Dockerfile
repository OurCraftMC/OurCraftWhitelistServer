FROM golang:1.21

WORKDIR /src
COPY . .

RUN go build -o /server

ENTRYPOINT ["/server"]
CMD ["/src/config_docker.yaml"]

