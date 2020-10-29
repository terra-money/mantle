FROM golang:1-buster as builder

WORKDIR /go/src/github.com/terra-project/mantle

COPY main.go main.go
COPY indexers indexers
COPY utils utils
COPY go.mod go.mod
COPY go.sum go.sum

RUN apt-get install git

RUN go mod download \
  && go build main.go

FROM ubuntu:latest

WORKDIR /app

COPY docker/libgo_cosmwasm.so /lib/libgo_cosmwasm.so
COPY --from=builder /go/src/github.com/terra-project/mantle/main /usr/local/bin/mantle

ENTRYPOINT [ "mantle" ]
