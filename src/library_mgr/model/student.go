package model


type Student struct {
	ID string `db:"id"`
	Name string `db:"name"`
	Gender string `db:gender`
	Grade string `db:grade`
	StudentID string `db:studentid`
	BorrowsBook []*Book
}
