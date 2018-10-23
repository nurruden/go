package main

import (
	"flag"
	"time"
	"fmt"
)

var period = flag.Duration("period",1*time.Second,"Sleep period")

func main(){
	flag.Parse()
	fmt.Printf("Sleeping for %v...",*period)
	time.Sleep(*period)
	fmt.Println()
}