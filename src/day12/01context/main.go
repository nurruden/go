package main

import (
	"fmt"
	"context"
)

func add(ctx context.Context,a,b int) int{
	traceld := ctx.Value("trace_id").(string)
	fmt.Printf("trace_id:%v\n",traceld)
	return a + b
}

func calc(ctx context.Context, a,b int) int{
	traceld := ctx.Value("trace_id").(string)
	fmt.Printf("trace_id:%v\n",traceld)
	return add(ctx,a,b)
}

func main(){
	ctx := context.WithValue(context.Background(),"trace_id","123456")
	calc(ctx,388,200)
}