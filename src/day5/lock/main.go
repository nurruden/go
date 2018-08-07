package main

import (
	"fmt"
	"sync"
)

var x = 0
func increament(wg *sync.WaitGroup,m *sync.Mutex){
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()

}

func main(){
	var w sync.WaitGroup
	var m sync.Mutex
	for i:=0;i<1000;i++{
		w.Add(1)
		go increament(&w,&m)
	}
	w.Wait()
	fmt.Printf("Final value of x: %d\n",x)
}