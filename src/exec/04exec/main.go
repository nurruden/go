package main

import (
	"os"
	"fmt"
)

func main(){

	w := os.Stdout
	w.Write([]byte("helloagain"))
	fmt.Println('\n')
	fmt.Printf("%T\n",w)
	//fmt.Println(w)

}
