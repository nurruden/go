package main

import (
	"bufio"
    "fmt"
    "os"
	
)
var (
	countEnglish = 0
	countSpace = 0
	countNumber = 0
	countSpecialCharacter = 0
)

func main(){
	fmt.Print("Please input character: ")
	reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
	for _, rune := range input {
        switch {
        case (rune >= 'A' && rune <= 'Z'):
            countEnglish ++
        case (rune >= 'a' && rune <= 'z'):
            countEnglish ++
        case rune ==  ' ' || rune == '\t':
            countSpace ++
        case (rune >= '0' && rune <= '9'):
	        countNumber ++
        default:
			countSpecialCharacter ++
        }
	}
	fmt.Printf("char=%d\nspace=%d\ndigit=%d\nothers=%d\n",countEnglish,countSpace,countNumber,countSpecialCharacter)
}