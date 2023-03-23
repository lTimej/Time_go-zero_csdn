FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /app

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .

RUN go build -ldflags="-s -w" -o /app/csdn/user/cmd/rpc/user_rpc ./csdn/user/cmd/rpc
RUN go build -ldflags="-s -w" -o /app/csdn/user/cmd/api/user_api ./csdn/user/cmd/api
RUN go build -ldflags="-s -w" -o /app/csdn/channel/cmd/rpc/article_rpc ./csdn/channel/cmd/rpc
RUN go build -ldflags="-s -w" -o /app/csdn/channel/cmd/rpc/article_api ./csdn/channel/cmd/api

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app /app/Time_go-zero_csdn

CMD ["./Time_go-zero_csdn/run.sh"]

