package main

import (
	"fmt"
	"time"
)

func testTimer(){
	// newtimer定时器，只触发一次
	timer := time.NewTimer(time.Second)
	for V := range timer.C {
		fmt.Printf("time:%v\n",V)
		// reset，否则会一直阻塞
		timer.Reset(time.Second)
}
}

func testTicker(){
	// newtimer定时器，每隔一秒触发一次
	timer := time.NewTicker(time.Second)
	for V := range timer.C {
		fmt.Printf("time:%v\n",V)
	
}
}


func timestampToTime(timestamp int64) {
	t := time.Unix(timestamp,0)
	fmt.Printf("time: %v\n",t)
}

func main() {
	// var now time.time
	now := time.Now()
	fmt.Printf("Current time is: %v\n",now)
	fmt.Printf("year:%d,month:%d,day:%d,hour:%2d, minute:%2d,second:%2d\n",
		now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
	fmt.Printf("timestamp:%d ns:%d\n",now.Unix(),now.UnixNano())

	// go testTimer()
	// testTicker()

	// time.Sleep(time.Minute)
	timestampToTime(now.Unix())
	
}