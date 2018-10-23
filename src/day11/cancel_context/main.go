package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//wg进行退出控制
var wg sync.WaitGroup

func worker(ctx context.Context) {
LOOP:
	for {
		fmt.Printf("worker\n")
		time.Sleep(time.Second)
		select { //select进行channel判断
		case <-ctx.Done():
			break LOOP
		default:

		}
	}
	wg.Done()

}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 3)
	cancel()
	wg.Wait()

}
