package utils

import (
	"fmt"
	"time"
)

func TimeToTimeStamp(t time.Time) int64 {
	return t.Unix()
}

func TimeStampToTimeString(t int64) string {
	return time.Unix(t, 0).Format("2006-01-02 15:04:05")
}

func TimeToString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func StringToTime(t string) time.Time {
	tt, _ := time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
	return tt
}

func TimeCompare(now, tt int64) string {
	var t string
	diff := now - tt
	t_year := time.Unix(tt, 0).Format("2006")
	n_year := time.Unix(now, 0).Format("2006")
	H := int64(time.Hour) / 1000000000
	M := int64(time.Minute) / 1000000000
	D := H * 24
	if diff < 60 {
		t = "刚刚"
	} else if diff > M && diff < H {
		t = fmt.Sprintf("%d分钟前", (diff)/60)
	} else if diff > H && diff < D {
		t = fmt.Sprintf("%d小时前", (diff)/3600)
	} else if diff > D && diff < 11*D {
		t = fmt.Sprintf("%d天前", (diff)/D)
	} else if diff > 10*D && t_year == n_year {
		t = time.Unix(tt, 0).Format("2006-01-02 15:04:05")[5:10]
	} else {
		t = time.Unix(tt, 0).Format("2006-01-02 15:04:05")
	}
	return t
}
