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

RUN go build -o /data/server/user-rpc  -v csdn/user/cmd/rpc/user.go
RUN go build -o /data/server/user-api  -v csdn/user/cmd/api/user.go
RUN go build -o /data/server/channel-rpc  -v csdn/channel/cmd/rpc/channel.go
RUN go build -o /data/server/channel-api  -v csdn/channel/cmd/api/api.go
RUN go build -o /data/server/im-rpc  -v csdn/im/cmd/rpc/im.go
RUN go build -o /data/server/im-api  -v csdn/im/cmd/api/api.go
RUN go build -o /data/server/product-rpc  -v csdn/shop_product/cmd/rpc/product.go
RUN go build -o /data/server/product-api  -v csdn/shop_product/cmd/api/api.go

COPY csdn/user/cmd/rpc/etc/user.yaml /etc/user/user.yaml
COPY csdn/user/cmd/api/etc/api-api.yaml /etc/user/api-api.yaml
COPY csdn/channel/cmd/rpc/etc/channel.yaml /etc/channel/channel.yaml
COPY csdn/channel/cmd/api/etc/api-api.yaml /etc/channel/api-api.yaml
COPY csdn/im/cmd/rpc/etc/im.yaml /etc/im/im.yaml
COPY csdn/im/cmd/api/etc/api-api.yaml /etc/im/api-api.yaml
COPY csdn/shop_product/cmd/rpc/etc/product.yaml /etc/product/product.yaml
COPY csdn/shop_product/cmd/api/etc/api-api.yaml /etc/product/api-api.yaml

COPY modd.conf /modd.conf
COPY modd /modd
RUN rm -rf /app/

FROM scratch

COPY --from=builder / /

CMD ["/modd"]

