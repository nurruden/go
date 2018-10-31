package main

import (
	"fmt"
	"sync"
	"time"
)

var a string
var once sync.Once

func setup(){
	a = "hello world"
	//time.Sleep(time.Second*5)
}

func doprint() {
	once.Do(setup)

	//setup()
	fmt.Println(a+"\n")

}

func twoprint(){
	go doprint()
	go doprint()
}

func main()  {
	twoprint()
	time.Sleep(time.Second*10)
}