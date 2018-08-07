package main

import (
	"fmt"
	"time"
)

func queryDB(ch chan int){
	time.Sleep(time.Second)
    ch <- 100
}

func main() {
	ch := make(chan int)
	go queryDB(ch)
	t := time.NewTicker(time.Second)
	select {
	// 在定时器时间内，100没有送给ch,ch就不能送给v,
	case v := <-ch:
		fmt.Println("result", v) 
	case <-t.C:
	    fmt.Println("timeout") }
}