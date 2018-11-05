package main

import "fmt"

func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		fmt.Printf("Current I is: %d\n",i)
		ch <- i
	}
}

func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// The prime sieve: Daisy-chain Filter processes.
func main() {
	ch := make(chan int) // Create a new channel.
	go Generate(ch)      // Launch Generate goroutine.
	for i := 0; i < 10; i++ {
		prime := <-ch
		//fmt.Println(prime)
		fmt.Printf("Prime is: %v\n",prime)
		ch1 := make(chan int)

		go Filter(ch, ch1, prime)
		ch = ch1
	}
}