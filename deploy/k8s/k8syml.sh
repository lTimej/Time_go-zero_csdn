#!/bin/bash

goctl kube deploy -replicas 2 -nodePort 31001 -requestCpu 200 -requestMem 50 -limitCpu 300 -limitMem 100 -name user-rpc -namespace csdn -image user_rpc:v1 -o yml/user/rpc.yml -port 1001

goctl kube deploy -replicas 2 -nodePort 32001 -requestCpu 200 -requestMem 50 -limitCpu 300 -limitMem 100 -name user-api -namespace csdn -image user_api:v1 -o yml/user/api.yml -port 2001

goctl kube deploy -replicas 2 -nodePort 31002 -requestCpu 200 -requestMem 50 -limitCpu 300 -limitMem 100 -name channel-rpc -namespace csdn -image channel_rpc:v1 -o yml/channel/rpc.yml -port 1002

goctl kube deploy -replicas 2 -nodePort 32002 -requestCpu 200 -requestMem 50 -limitCpu 300 -limitMem 100 -name channel-api -namespace csdn -image channel_api:v1 -o yml/channel/api.yml -port 2002

goctl kube deploy -replicas 2 -nodePort 31003 -requestCpu 200 -requestMem 50 -limitCpu 300 -limitMem 100 -name im-rpc -namespace csdn -image im_rpc:v1 -o yml/im/rpc.yml -port 1003

goctl kube deploy -replicas 2 -nodePort 32003 -requestCpu 200 -requestMem 50 -limitCpu 300 -limitMem 100 -name im-api -namespace csdn -image im_api:v1 -o yml/im/api.yml -port 2003

goctl kube deploy -replicas 2 -nodePort 31004 -requestCpu 200 -requestMem 50 -limitCpu 300 -limitMem 100 -name product-rpc -namespace csdn -image product_rpc:v1 -o yml/product/rpc.yml -port 1004

goctl kube deploy -replicas 2 -nodePort 32004 -requestCpu 200 -requestMem 50 -limitCpu 300 -limitMem 100 -name product-api -namespace csdn -image product_api:v1 -o yml/product/api.yml -port 2004

goctl kube deploy -replicas 2 -nodePort 31005 -requestCpu 200 -requestMem 50 -limitCpu 300 -limitMem 100 -name order-rpc -namespace csdn -image order_rpc:v1 -o yml/order/rpc.yml -port 1005

goctl kube deploy -replicas 2 -nodePort 32005 -requestCpu 200 -requestMem 50 -limitCpu 300 -limitMem 100 -name order-api -namespace csdn -image order_api:v1 -o yml/order/api.yml -port 2005

goctl kube deploy -replicas 2 -nodePort 31006 -requestCpu 200 -requestMem 50 -limitCpu 300 -limitMem 100 -name scheduler -namespace csdn -image scheduler:v1 -o yml/scheduler/scheduler.yml -port 1006

goctl kube deploy -replicas 2 -nodePort 32006 -requestCpu 200 -requestMem 50 -limitCpu 300 -limitMem 100 -name job -namespace csdn -image job:v1 -o yml/scheduler/job.yml -port 2006

goctl kube deploy-replicas 2 -nodePort 32007 -requestCpu 200 -requestMem 50 -limitCpu 300 -limitMem 100 -name city-api -namespace csdn -image city_api:v1 -o yml/city/api.yml -port 2005