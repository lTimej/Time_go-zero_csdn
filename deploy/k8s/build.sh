#!/bin/bash

docker build -t user_rpc:v1 -f user_rpc.Dockerfile .
docker build -t user_api:v1 -f user_api.Dockerfile .

docker build -t channel_rpc:v1 -f channel_rpc.Dockerfile .
docker build -t channel_api:v1 -f channel_api.Dockerfile .

docker build -t im_rpc:v1 -f im_rpc.Dockerfile .
docker build -t im_api:v1 -f im_api.Dockerfile .

docker build -t product_rpc:v1 -f product_rpc.Dockerfile .
docker build -t product_api:v1 -f product_api.Dockerfile .

docker build -t order_rpc:v1 -f order_rpc.Dockerfile .
docker build -t order_api:v1 -f order_api.Dockerfile .

docker build -t scheduler:v1 -f scheduler.Dockerfile .
docker build -t job:v1 -f job.Dockerfile .

docker build -t city_api:v1 -f city_api.Dockerfile .