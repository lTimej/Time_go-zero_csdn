package utils

import (
	"liujun/Time_go-zero_csdn/common/snowflak"
)

func GetSN(prefix string) string {
	WI, _ := snowflak.NewSnowFlak(1, 2, 0, -1)
	user_id := WI.GetId()
	random_str := GetRandstring(3)
	return prefix + random_str + user_id
}
