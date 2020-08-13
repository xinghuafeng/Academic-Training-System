package util

import "time"

const (
	SecondsInWeek = 3600 * 24 * 7 // 一周换算成秒

	TimeLayout = "2006-01-02 15:04:05" // 据说是 golang 诞生时间
)

/*
	格式化当前系统时间
*/
func CurrentTime() string {
	timeStr := time.Now().Format(TimeLayout)
	return timeStr
}
