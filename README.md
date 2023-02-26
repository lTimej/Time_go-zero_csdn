# Time_go-zero_csdn
this is a go-zero exercise
# 快速生成rpc服务
goctl rpc new 服务名
# 快速生成api
goctl api new api
# 通过api文件生成go文件
goctl api go -api core.api -dir . -style go_zero
# 通过proto生成go文件
goctl rpc protoc user.proto --go_out=. --go-grpc_out=. --zrpc_out=.