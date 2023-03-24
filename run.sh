
/app/Time_go-zero_csdn/csdn/user/cmd/rpc/user_rpc -f /app/Time_go-zero_csdn/csdn/user/cmd/rpc/etc/user.yaml&

/app/Time_go-zero_csdn/csdn/channel/cmd/rpc/article_rpc -f /app/Time_go-zero_csdn/csdn/channel/cmd/rpc/etc/channel.yaml&

sleep 10

/app/Time_go-zero_csdn/csdn/user/cmd/api/user_api -f /app/Time_go-zero_csdn/csdn/user/cmd/api/etc/api-api.yaml&

/app/Time_go-zero_csdn/csdn/channel/cmd/api/article_api -f /app/Time_go-zero_csdn/csdn/channel/cmd/api/etc/api-api.yaml
