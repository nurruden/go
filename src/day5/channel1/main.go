package main

import (
	"time"
	"fmt"
)

func main(){
	var intchan chan int = make(chan int,1)
	go func(){
		intchan <- 100
		fmt.Print("Insert end\n")
	}()
	go func(){
		fmt.Printf("start\n")
		time.Sleep(time.Second*3)
		var a int
		a = <- intchan
		fmt.Printf("a=%d\n",a)
	}()
	time.Sleep(time.Second*3)
	
}