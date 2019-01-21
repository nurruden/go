package main

import (
	"fmt"
	"time"
)

const (
	  A = 'a'
      B = iota
      C
      D = "NICK"
      E
	  H = iota
	  I
      F = 618
      G

)

func sw3() {
	switch num := 5; {
	case num > 0 && num < 4:
		fmt.Printf("%v in 1,2,3\n", num)
	case num > 4 && num < 7:
		fmt.Printf("%v in 4,5,6\n", num)
		fallthrough
	case num > 7 && num < 10:
		fmt.Printf("%v big 789\n", num)
	default:
		fmt.Printf("default...\n")
	}
}

func lab() {
LABLE:
	for i := 0; i < 10; i++ {
		for true {
			i++
			if i == 6 {
				break LABLE
			}
			fmt.Println(i)
		}
	}
}

func got() {
	i := 0
HERE:
	fmt.Println(i)
	i++
	if i == 5 {
		return
	}
	goto HERE
}


func main()  {
	//fmt.Printf("A:%v\n,B:%v\n,C:%v\n,D:%v\n,E:%v\n,F:%v\n,G:%v\n,H:%v\n,I:%v\n",A,B,C,D,E,F,G,H,I)
    //sw3()
	//lab()
	//got()
	//ch := make(chan int, 1)
	//for {
	//	select {
	//	case ch <- 0:
	//	case ch <- 1:
	//	}
	//	i := <-ch
	//	fmt.Println(i)
	//	time.Sleep(time.Second)
	//}
	//ch := make(chan bool)
	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(time.Second*2)
		timeout <- true
	}()

	select {
	//case <- ch:
	//	fmt.Println("This is ch.")
	case <- timeout:
		fmt.Println("This is timeout.")
	}
	//var ptr *int
	num := 100
	ptr := &num
	fmt.Println(ptr)    //0xc42000e1f8
	fmt.Println(*ptr)   //100
	*ptr = 200
	fmt.Println(num)    //200

}