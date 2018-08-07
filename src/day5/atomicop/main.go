package main

import (
	"sync/atomic"
	"time"
	"fmt"
)

var count int32 
func test1(){
	for i := 0; i < 1000000;i++{
		atomic.AddInt32(&count,1)
	}
}

func test2(){
	for i := 0; i < 1000000;i++{
	    atomic.AddInt32(&count,1)
	}
}

func main(){
	go test1()
	go test2()

	time.Sleep(time.Second)
	fmt.Printf("%d\n",count)
}