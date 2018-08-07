package main

import (
	"fmt"
	"time"
)

func main(){
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	time.Sleep(5*time.Second)
}
func sendData(ch chan string) {
	ch <- "Beijing"
	ch <- "Shanghai"
	ch <- "Nanjing"
	ch <- "Haerbin"
	ch <- "London"
	// fmt.Println(ch)
}

func getData(ch chan string){
	var inChan string
	for {
		inChan = <- ch
		fmt.Println(inChan)
	}

	close(ch)
	time.Sleep(time.Second*10)
	for i:=0;i<10;i++{
		<- ch
	}
}
