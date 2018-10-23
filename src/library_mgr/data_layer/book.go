package data_layer

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"library_mgr/model"
)

func InsertBook(book *model.Book) (err error) {
	if book == nil {
		err = fmt.Errorf("invalid book parameter")
		return
	}

	sqlstr := "select bookid from book where bookid=?"
	var bookId string
	err = Db.Get(&bookId, sqlstr, book.BookId)
	if err == sql.ErrNoRows {
		//插入操作
		sqlstr = `insert into book(
					  author, bookname, publishtime, stocknum, bookid
				  )values(?, ?, ?, ?, ?)`
		_, err = Db.Exec(sqlstr, book.Author, book.BookName, book.PublishTime, book.StockNum, book.BookId)
		if err != nil {
			return
		}
		return
	}

	if err != nil {
		return
	}

	err = fmt.Errorf("book_id:%s is already exists", bookId)
	return
}

func UpdateBook(book *model.Book) (err error) {
	if book == nil {
		err = fmt.Errorf("invalid book parameter")
		return
	}

	//插入操作
	sqlstr := `update book set
						author = ?, bookname=?, publishtime=?,  stocknum= stocknum+?
				  where 
				        bookid = ?`
	result, err := Db.Exec(sqlstr, book.Author, book.BookName, book.PublishTime, book.StockNum, book.BookId)
	if err != nil {
		return
	}

	affects, err := result.RowsAffected()
	if err != nil {
		return
	}

	if affects == 0 {
		err = fmt.Errorf("update book failed, book_id:%s, not found", book.BookId)
		return
	}
	return
}

func QueryBook(bookId string) (book *model.Book, err error) {
	//插入操作
	sqlstr := `select 
					author , bookname, publishtime,  stocknum, bookid, id
				from book
				where 
					bookid = ?`
	book = &model.Book{}
	err = Db.Get(book, sqlstr, bookId)
	if err != nil {
		return
	}
	return
}

