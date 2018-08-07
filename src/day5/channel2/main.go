package main

import ( 
	"fmt"
	"time" 
)
func main() {
	ch := make(chan string,100)
	go sendData(ch)
	time.Sleep(100 * time.Second)
}
func sendData(ch chan string) { 
	var i int
    for {
		var str string
		str = fmt.Sprintf("stu %d",i)
		fmt.Println("write:",str)
		ch <- str
		i++
	} 
}