package main

import "fmt"

func main(){
    var (
        a = 10
        f = 10.0
        str string = "hello"
    )
    fmt.Printf("%d\n", a)
    fmt.Printf("%s\n",str)
    fmt.Printf("%b\n",a)
    fmt.Printf("%x\n",a)
    fmt.Printf("%f\n",f)
}
