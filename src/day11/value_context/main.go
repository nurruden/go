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
	traceCode, ok := ctx.Value("TRACE_CODE").(string)
	if ok {
		fmt.Printf("trace_code : %s \n", traceCode)
	}
LOOP:
	for {
		fmt.Printf("worker,trace_code : %s\n", traceCode)
		time.Sleep(time.Second)
		select { //select进行channel判断
		case <-ctx.Done():
			break LOOP
		default:

		}
	}
	fmt.Printf("worker done ,traceCode :%s \n", traceCode)
	wg.Done()

}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*500)
	ctx = context.WithValue(ctx, "TRACE_CODE", "23123123123")
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 3)
	cancel()
	wg.Wait()

}
