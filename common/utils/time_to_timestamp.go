package utils

import "time"

func TimeToTimeStamp(t time.Time) int64 {
	return t.Unix()
}

func TimeStampToTimeString(t int64) string {
	return time.Unix(t, 0).Format("2006-01-02 15:04:05")
}

func TimeToString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
