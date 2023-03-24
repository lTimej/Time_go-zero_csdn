FROM golang:1.19 AS builder

LABEL stage=gobuilder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64


WORKDIR /app

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .

RUN go build -o /app/csdn/user/cmd/rpc/user_rpc /app/csdn/user/cmd/rpc
RUN go build -o /app/csdn/user/cmd/api/user_api /app/csdn/user/cmd/api
RUN go build -o /app/csdn/channel/cmd/rpc/article_rpc /app/csdn/channel/cmd/rpc
RUN go build -o /app/csdn/channel/cmd/api/article_api /app/csdn/channel/cmd/api

RUN chmod +x /app/csdn/user/cmd/rpc/user_rpc
RUN chmod +x /app/csdn/user/cmd/api/user_api
RUN chmod +x /app/csdn/channel/cmd/rpc/article_rpc
RUN chmod +x /app/csdn/channel/cmd/api/article_api

FROM debian:stretch-slim


WORKDIR /app
COPY --from=builder /app /app/Time_go-zero_csdn

#ENTRYPOINT ["./Time_go-zero_csdn/run.sh"]

