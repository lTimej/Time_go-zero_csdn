package utils

import "time"

func TimeToTimeStamp(t time.Time) int64 {
	return t.UnixMilli()
}
