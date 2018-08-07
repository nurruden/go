package main

import (
	"fmt"
	// "time"
)

func main() {
	ch := make(chan string)
	exitChan := make(chan bool,3)
	go sendData(ch,exitChan)
	go getData(ch,exitChan)
	go getData2(ch,exitChan)
	<- exitChan
	<- exitChan
	<- exitChan
}

func sendData(ch chan string,exitCh chan bool) {
	ch <- "aaa"
	ch <- "bbb"
	ch <- "ccc"
	ch <- "ddd"
	ch <- "eee"
	close(ch)
	exitCh <- true
}

func getData (ch chan string,exitCh chan bool) {
	
	for {
		input,ok := <- ch
		if !ok {
			break
		}
		fmt.Printf("getData中的input值:%s\n", input)
	}
	fmt.Printf("get data exited\n")
	exitCh <- true
}

func getData2 (ch chan string,exitCh chan bool) {
	
	for {
		input2,ok := <- ch
	    if !ok{
			break
		}
		fmt.Printf("getData2中的input值:%s\n", input2)
	}
	fmt.Printf("get data2 exited\n")
	exitCh <- true
}