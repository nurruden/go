package main

import (
	"fmt"
	"sync"
)
// func Count(ch chan int){
// 	fmt.Println("counting")
// 	ch <- 1
// }


// func main() {
// 	chs := make([]chan int,10)
// 	for i:=0;i<10;i++ {
// 		chs[i] = make(chan int)
// 		go Count(chs[i])
// 	}

// 	for _,ch :=range(chs) {
// 		<-ch
// 	}
// }

// func main(){
// 	ch :=make(chan int,1)
// 	for {
// 		select {
// 			case ch <- 0:
// 			case ch <- 1:
// 		}
// 		i := <-ch
// 		fmt.Println("Value received:",i)
// 	}
// }
var a string
var once sync.Once

func setup() {
	a = "Hello"
}

func doprint(){
	once.Do(setup)
	fmt.Print(a)
}

func twoprint() {
	go doprint()
	go doprint()
}

func main(){
	twoprint()
}