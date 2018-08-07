package bookinfo

import (
	"fmt"
	"os"
	"encoding/json"
	"log"
)

var flagbook bool = false

type Student struct {
	ID int `json:"ID"`
	Username string `json:"username"`
}

type Book struct {
	BookID int `json:"bookid"`
	BookName string `json:"bookname"`
	BookCopy int `json:"bookcopy"`
	BookAuthor string `json:"bookauthor"`
	PublishTime string `json:"publishtime"`
	StudentArray []Student `json:"studentarray"`

}



var book []Book = make([]Book,0,0)

func (pbook *Book)AddBookInfo() {

	fmt.Println("Please input book ID:")
	fmt.Scanf("%d\n",&pbook.BookID)
	fmt.Println("Please input bookt name:")
	fmt.Scanf("%s\n",&pbook.BookName)
	fmt.Println("Please input book copy:")
	fmt.Scanf("%d\n",&pbook.BookCopy)
	fmt.Println("Please input book author:")
	fmt.Scanf("%s\n",&pbook.BookAuthor)
	fmt.Println("Please input book publish time:")
	fmt.Scanf("%s\n",&pbook.PublishTime)
    fmt.Print("BookID\tBookName\tBookCopy\tBookAuthor\tPublishTime\n")
	fmt.Printf("%v\t%v\t\t%v\t\t%v\t\t%v\n",pbook.BookID,pbook.BookName,pbook.BookCopy,pbook.BookAuthor,pbook.PublishTime)
	book = append(book,*pbook)
	bookdata, err := json.Marshal(book)
	if err != nil{
		fmt.Printf("%v\n",err)
	}
	fp, err := os.OpenFile("bookdata.json", os.O_RDWR|os.O_CREATE, 0755)
    if err != nil {
        log.Fatal(err)
    }
    defer fp.Close()
    _, err = fp.Write(bookdata)
    if err != nil {
        log.Fatal(err)
    }
}

func readBookinfo() {
	fp, err := os.OpenFile("./bookdata.json", os.O_RDONLY, 0755)
    defer fp.Close()
    if err != nil {
        log.Fatal(err)
    }
    data := make([]byte, 1000)
    n, err := fp.Read(data)
    if err != nil {
        log.Fatal(err)
	}
	var bookslice []Book
	if err := json.Unmarshal(data[:n],&bookslice);err != nil{
		fmt.Printf("%v\n",err)
	}
    for _,v := range(bookslice){
		fmt.Printf("%v\t%v\t\t%v\t\t%v\t%v\n",v.BookID,v.BookName,v.BookCopy,v.BookAuthor,v.PublishTime)
	}
}

func initBook(){

	var selectStr string

	fmt.Print("\t\t\t\tWelcome to the book manage system\t\t\t\n")
	fmt.Print("**********************************************************************************************************\n")
	fmt.Print("Please select the NO:\n")
	fmt.Print("a)Add new book information\nb)Display book info\nc)Quit\n")
	fmt.Scanf("%s\n", &selectStr)

	switch selectStr{
	case "a":
		fmt.Printf("Now add new book: \n")
		var pbook *Book = new(Book)
		pbook.AddBookInfo()
	
	case "b":
		fmt.Print("Currently the student info are:\n")
		fmt.Print("BookID\tBookName\tBookCopy\tBookAuthor\tPublishTime\n")
		readBookinfo()
	
		break
	
	case "c":
		flagbook= true
		break
	default:
		break
	}

}
func RunBook(){
	for {
		initBook()
		if flagbook {
			break
		}
	}
}