#!/bin/bash

cd csdn/user/cmd/rpc/
goctl docker -go user.go
cd ../../../../
mv csdn/user/cmd/rpc/Dockerfile /home/time/Time_go-zero_csdn/user_rpc.Dockerfile

cd csdn/user/cmd/api/
goctl docker -go user.go
cd ../../../../
mv csdn/user/cmd/api/Dockerfile /home/time/Time_go-zero_csdn/user_api.Dockerfile

cd csdn/channel/cmd/rpc/
goctl docker -go channel.go
cd ../../../../
mv csdn/channel/cmd/rpc/Dockerfile /home/time/Time_go-zero_csdn/channel_rpc.Dockerfile

cd csdn/channel/cmd/api/
goctl docker -go api.go
cd ../../../../
mv csdn/channel/cmd/api/Dockerfile /home/time/Time_go-zero_csdn/channel_api.Dockerfile

cd csdn/im/cmd/rpc/
goctl docker -go im.go
cd ../../../../
mv csdn/im/cmd/rpc/Dockerfile /home/time/Time_go-zero_csdn/im_rpc.Dockerfile

cd csdn/im/cmd/api/
goctl docker -go api.go
cd ../../../../
mv csdn/im/cmd/api/Dockerfile /home/time/Time_go-zero_csdn/im_api.Dockerfile

cd csdn/shop_product/cmd/rpc/
goctl docker -go product.go
cd ../../../../
mv csdn/shop_product/cmd/rpc/Dockerfile /home/time/Time_go-zero_csdn/product_rpc.Dockerfile

cd csdn/shop_product/cmd/api/
goctl docker -go api.go
cd ../../../../
mv csdn/shop_product/cmd/api/Dockerfile /home/time/Time_go-zero_csdn/product_api.Dockerfile

cd csdn/order/cmd/rpc/
goctl docker -go order.go
cd ../../../../
mv csdn/order/cmd/rpc/Dockerfile /home/time/Time_go-zero_csdn/order_rpc.Dockerfile

cd csdn/order/cmd/api/
goctl docker -go api.go
cd ../../../../
mv csdn/order/cmd/api/Dockerfile /home/time/Time_go-zero_csdn/order_api.Dockerfile

cd csdn/scheduler/cmd/scheduler/
goctl docker -go scheduler.go
cd ../../../../
mv csdn/scheduler/cmd/scheduler/Dockerfile /home/time/Time_go-zero_csdn/scheduler.Dockerfile

cd csdn/scheduler/cmd/job/
goctl docker -go job.go
cd ../../../../
mv csdn/scheduler/cmd/job/Dockerfile /home/time/Time_go-zero_csdn/job.Dockerfile

cd csdn/city/cmd/api/
goctl docker -go api.go
cd ../../../../
mv csdn/city/cmd/api/Dockerfile /home/time/Time_go-zero_csdn/city_api.Dockerfile