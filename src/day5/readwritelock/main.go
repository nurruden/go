package main

import (
	"time"
	"fmt"
	"sync"
)

var rwlock sync.RWMutex
var wg sync.WaitGroup
var count int
func writer(){
	for i := 0;i < 1000;i++{
		rwlock.Lock()
		count++
		time.Sleep(10*time.Millisecond)
		rwlock.Unlock()
	}
    wg.Done()
}

func reader(){
    for i := 0;i < 1000;i++{
		rwlock.RLock()
		fmt.Printf("count=%d\n",count)
		rwlock.RUnlock()
	}
	wg.Done()
}

func main(){
	wg.Add(2)
	go writer()
	go reader()
	// for i :=0;i<10;i++{
	// 	wg.Add(1)
	// 	go reader()
	// }
	wg.Wait()
}