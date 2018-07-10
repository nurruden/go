package main

import (
	"fmt"
	"time"
)

func main() {
	// var now time.time
	now := time.Now()
	fmt.Printf("Current time is: %v\n",now)
	fmt.Printf("year:%d,month:%d,day:%d,hour:%2d, minute:%2d,second:%2d\n",
		now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
	// 必须是"02/01/2006 15:04:05"这个时间
	str := now.Format("02/01/2006 15:04:05")
	fmt.Printf("str:%s\n",str)
}