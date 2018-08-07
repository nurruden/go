package main

import (
	"fmt"
	"day4/homework/student"
	"day4/homework/book"    
)

var flag bool = false

func run() {

	var selectNum int

	fmt.Print("\t\t\t\tWelcome to the book manage system\t\t\t\n")
	fmt.Print("#######################################################################################################################\n")
	fmt.Print("Please select the NO:\n")
	fmt.Print("1)Book manage system\n2)Student manage system\n3)Borrow a book\n4)Exit\n")
	fmt.Scanf("%d\n", &selectNum)

	switch {
	case selectNum == 2:
		fmt.Printf("Welcome to the student manage system \n")
		stuinfo.RunStudent()
	case selectNum == 1:
		fmt.Printf("Welcome to the book manage system\n")
        bookinfo.RunBook()
	case selectNum == 3:
		fmt.Print("Welcome to the book borrowing manage system\n")
		break
	case selectNum == 4:
		flag = true
		break
	default:
		break
	}

}


func main(){
	
	for {
		run()
		
		if flag {
			break
		}
	}
}