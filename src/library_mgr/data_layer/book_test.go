package data_layer

import (
	"library_mgr/model"
	"testing"
	"time"
)

func init() {
	dns := "root:Embe1mpls@tcp(127.0.0.1:3306)/bookmgr?parseTime=True"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

func TestInsertBook(t *testing.T) {
	var book model.Book
	book.Author = "allan"
	book.BookId = "123qwe45"
	book.BookName = "python"
	book.PublishTime = time.Now()
	book.StockNum = 10

	err := InsertBook(&book)
	if err != nil {
		t.Errorf("insert book failed, err:%v", err)
		return
	}

	t.Logf("insert book succ")
}

func TestUpdateBook(t *testing.T) {
	var book model.Book
	book.Author = "allan"
	book.BookId = "123qwe45"
	book.BookName = "python"
	book.PublishTime = time.Now()
	book.StockNum = 10

	err := UpdateBook(&book)
	if err != nil {
		t.Errorf("update book failed, err:%v", err)
		return
	}

	t.Logf("update book succ")
}


func TestQueryBook(t *testing.T) {
	var book *model.Book
	bookId := "123qwe45"
	book, err := QueryBook(bookId)
	if err != nil {
		t.Errorf("query book failed, err:%v", err)
		return
	}

	t.Logf("query book succ, book:%#v", book)
}
