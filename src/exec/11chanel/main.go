package main

import "fmt"

//func main() {
//	messages := make(chan string)
//	go func() {
//		messages <- "ping"
//	}()
//	msg := <-messages
//	fmt.Println(msg)
//}

func main() {
	messages := make(chan string,1)

	messages <- "ping"
	msg := <-messages


	fmt.Println(msg)
	//select {
	//case messages <- "msg":
	//	fmt.Println("sent message")
	//default:
	//	fmt.Println("no message sent")
	//}
}

//func main(){
//	queue := make(chan string, 2)
//	queue <- "one"
//	queue <- "two"
//	close(queue)
//	fmt.Println(len(queue))
//	//go func() {
//	//	fmt.Println(queue)
//	//	for elem := range queue {
//	//		fmt.Println(elem)
//	//	}
//	//}()
//
//	for elem := range queue {
//		fmt.Println(elem)
//	}
//}