package main

import (
	"fmt"
	"math/rand"
)

func rand_generator_2() chan int {
	out := make(chan int)
	go func() {
		for {
			out <- rand.Int()
		}
	}()
	return out
}

func rand_generator_3() chan int {
	rand_generato_1 := rand_generator_2()
	rand_generato_2 := rand_generator_2()
	out := make(chan int)

	go func() {
		for {
			//读取生成器1中的数据，整合
			out <- <-rand_generato_1
		}
	}()
	go func() {
		for {
			//读取生成器2中的数据，整合
			out <- <-rand_generato_2
		}
	}()
	return out
}

func main() {
	// 生成随机数作为一个服务
	rand_service_handler := rand_generator_3()
	fmt.Printf("%d\n", <-rand_service_handler)
}