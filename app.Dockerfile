FROM golang:1.19 AS builder

LABEL stage=gobuilder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /csdn
ADD go.mod .
ADD go.sum .
RUN go mod download
