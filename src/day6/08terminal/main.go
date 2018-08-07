package main

import (
	"os"
	"fmt"
	// "io"
	"bufio"
)

func main(){
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString("hello world\n")
	writer.Flush()

	reader := bufio.NewReader(os.Stdin)
	data,err := reader.ReadString('\n')

	if err != nil {
		fmt.Printf("read from console failed:%v\n",err)
		return
	}
	fmt.Printf("DONE!%v\n",data)
}