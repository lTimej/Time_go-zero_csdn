# Time_go-zero_csdn
this is a go-zero exercise
# 方式一
## 快速生成rpc和api
goctl rpc new rpc
goctl api new api
# 方式二
## 通过api文件生成go文件
goctl api go -api user.api -dir . -style go_zero
goctl api go -api channel.api -dir . -style go_zero
goctl api go -api im.api -dir . -style go_zero
goctl api go -api product.api -dir . -style go_zero
goctl api go -api order.api -dir . -style go_zero
## 通过proto生成go文件
goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
goctl rpc protoc channel.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
goctl rpc protoc im.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
goctl rpc protoc product.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
goctl rpc protoc order.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.

# 通过ddl生成model
goctl model mysql ddl -src="./sql/user.sql" -dir="./csdn/user/model" -c
goctl model mysql ddl -src="./sql/article.sql" -dir="./csdn/channel/model" -c
goctl model mysql ddl -src="./sql/im.sql" -dir="./csdn/im/model" -c
goctl model mysql ddl -src="./sql/shop_product.sql" -dir="./csdn/shop_product/model" -c
goctl model mysql ddl -src="./sql/order.sql" -dir="./csdn/order/model" -c
goctl model mysql ddl -src="./sql/city.sql" -dir="./csdn/city/model" -c

# 单体服务
goctl api go -api city.api -dir . -style go_zero

### user.proto
syntax = "proto3";

package user;
option go_package = "./user";

message IdRequest {
string id = 1;
}

message UserResponse {
// 用户id
string id = 1;
// 用户名称
string name = 2;
// 用户性别
string gender = 3;
}

service User {
rpc getUser(IdRequest) returns(UserResponse);
}

### user.api
type Request {
Name string `path:"name,options=you|me"`
}

type Response {
Message string `json:"message"`
}

service api-api {
@handler ApiHandler
get /from/:name(Request) returns (Response)
