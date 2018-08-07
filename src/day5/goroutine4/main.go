package main

import (
	"fmt"
	"time"
)

func main() {
    c := make(chan string)
    go func(){
        c <- "hello"
    }()
    
    go func(){
       word := <- c + " world"
       fmt.Println(word)
    }()
    time.Sleep(1 * time.Second)
}
