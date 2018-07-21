package main

import (
	"fmt"
)

func main() {
    var a, b int
    for a = 2; a <= 100; a++ {
        for b = 2; b <= (a / b); b++ {
            if a%b == 0 {
                break
            }
        }
        if b > (a / b) {
            fmt.Printf("%d\t是素数\n", a)
        }
    }
}