package borrowbook

import (
	"fmt"
)

var flagborrowbook bool = false
func initBook(){

	var selectStr string

	fmt.Print("\t\t\t\tWelcome to the book borrow system\t\t\t\n")
	fmt.Print("**********************************************************************************************************\n")
	fmt.Print("Please select the NO:\n")
	fmt.Print("a)Search a book\nb)Check student status\nc)Borrow a book\nd)Quit")
	fmt.Scanf("%s\n", &selectStr)

	switch selectStr{
	case "a":
		fmt.Printf("Search a book: \n")
		
	
	case "b":
		fmt.Print("Check student status:\n")
	
		break
	case "c":
		fmt.Print("Borrow a book:")
	case "d":
		flagborrowbook= true
		break
	default:
		break
	}

}
func RunBook(){
	for {
		initBook()
		if flagborrowbook {
			break
		}
	}
}