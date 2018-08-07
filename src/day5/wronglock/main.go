package main

import (
	"fmt"
	"sync"
)

var x = 0
func increament(wg *sync.WaitGroup){

	x = x + 1

	wg.Done()

}

func main(){
	var w sync.WaitGroup

	for i:=0;i<1000;i++{
		w.Add(1)
		go increament(&w)
	}
	w.Wait()
	fmt.Printf("Final value of x: %d\n",x)
}