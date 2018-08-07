package main

import (
	"sync"
	"fmt"
	// "time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)
	// exitChan := make(chan bool,3)
	wg.Add(3)
	go sendData(ch,&wg)
	go getData(ch,&wg)
	go getData2(ch,&wg)
	wg.Wait()
    
}

func sendData(ch chan string,waitGroup *sync.WaitGroup) {
	ch <- "aaa"
	ch <- "bbb"
	ch <- "ccc"
	ch <- "ddd"
	ch <- "eee"
	close(ch)
	fmt.Print("Send data exited\n")
	waitGroup.Done()
}

func getData (ch chan string,waitGroup *sync.WaitGroup) {
	
	for {
		input,ok := <- ch
		if !ok {
			break
		}
		fmt.Printf("getData中的input值:%s\n", input)
	}
	fmt.Printf("get data exited\n")
	waitGroup.Done()
}

func getData2 (ch chan string,waitGroup *sync.WaitGroup) {
	for v := range ch {
		fmt.Printf("Get data2 %s\n",v)
	}
	// for {
	// 	input2,ok := <- ch
	//     if !ok{
	// 		break
	// 	}
	// 	fmt.Printf("getData2中的input值:%s\n", input2)
	// }
	fmt.Printf("get data2 exited\n")
	waitGroup.Done()
}