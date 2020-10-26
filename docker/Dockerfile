FROM golang:alpine as builder

WORKDIR /app

COPY main.go ./
COPY go.* ./

RUN apk add --no-cache git build-base

RUN GOSUMDB=off go mod download \
  && go build main.go
