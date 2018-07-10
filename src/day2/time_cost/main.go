package main

import (
	"time"
	"fmt"
)

func test() {
	for i := 0; i < 1000000000;i++ {
		_=i
	}
}

func main() {
	start := time.Now().UnixNano()
	test()
	end := time.Now().UnixNano()

	fmt.Printf("cost=%d us\n",(end - start)/1000)
}