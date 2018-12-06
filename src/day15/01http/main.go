package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main(){
	conn,err := net.Dial("tcp","localhost:50000")
	if err != nil {
		fmt.Println("Error dialing",err.Error())
		return
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input,_ := inputReader.ReadString('\n')
		trimmedlnput := strings.Trim(input,"\r\n")
		if trimmedlnput == "Q"{
			return
		}
		_,err = conn.Write([]byte(trimmedlnput))
		if err != nil {
			return
		}
	}
}
