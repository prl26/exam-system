package utils

import (
	"fmt"
	"time"
)

func StringToTime(string2 string) time.Time {
	var layout = "2006-01-02 15:04:05" //转换的时间字符串带秒则 为 2006-01-02 15:04:05
	timeVal, errByTimeConver := time.ParseInLocation(layout, string2, time.Local)
	if errByTimeConver != nil {
		fmt.Println("TimeStr To Time Error.....", errByTimeConver)
	}
	return timeVal
}
