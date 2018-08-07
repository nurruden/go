package main

import (
	"time"
	"fmt"
	"runtime"
)

func main(){
	cpu := runtime.NumCPU()
	// runtime.GOMAXPROCS()

	for i :=0;i<8;i++{
		go func(){
			for {

			}
		}()
	}
	fmt.Printf("%d\n",cpu)
	time.Sleep(time.Second*5)
}