package model

import "time"

type Book struct {
	Id int64 `db:"id"`
	Author string `db:"author"`
	BookName string 	`db:"bookname"`
	PublishTime time.Time `db:"publishtime"`
	StockNum uint `db:"stocknum"`
	BookId  string `db:"bookid"`
}
