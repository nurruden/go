package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

//func QueryRow(Db *sqlx.DB) {
//	id := 1
//	var user User
//	err := Db.Get(&user, "select id,name,age from user where id>=?", id)
//	if err == sql.ErrNoRows {
//		fmt.Printf("no record found\n")
//		return
//	}
//	if err != nil {
//		fmt.Printf("get faild ,err:%v\n", err)
//		return
//	}
//
//	fmt.Printf("get user succ  user: %#v\n", user)
//}
//
func Query(Db *sqlx.DB) {
	var user []*User
	id := 0
	err := Db.Select(&user,"select id,name,age from user where id>?",id)
	if err != nil{
		return
	}
	fmt.Printf("user: %#\n",user)
}

func Insert(Db *sqlx.DB) {
	username := "user01"
	age := 18

	result,err := Db.Exec("insert into user(name, age)values(?,?)",username,age)
	if err != nil{
		fmt.Printf("ecec failed,err:%v\n",err)
		return
	}

	id,err := result.LastInsertId()
	if err != nil {
		fmt.Printf("last insert id failed,err:%v\n",id,err)
		return
	}

	affectRows,err := result.RowsAffected()
	if err != nil {
		fmt.Printf("offefectRows failed,err:%v\n",err)
	}

	fmt.Printf("last insert id:%d affect rows:%d\n",id,affectRows)
}

//func QuerySelect(Db *sqlx.DB) {
//	var user []*User //存储的是user对象的指针,即地址
//	id := 0
//	//切片是引用类型,传入切片的地址,是因为切片可能会扩容,导致地址变化,相当于切片接收append
//	err := Db.Select(&user, "select id,name,age from user where id >?", id)
//	if err != nil {
//		return
//	}
//	fmt.Printf("user :%#v\n", user)
//
//}

func main() {
	dsn := "root:Embe1mpls@tcp(127.0.0.1:3306)/gotest"
	//dsn := "root:123456@tcp(192.168.10.200:3306)/golang"
	Db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("open mysql faild,err %v\n", err)
		return
	}
	err = Db.Ping()
	if err != nil {
		fmt.Printf("ping  faild,err %v\n", err)
		return
	}

	fmt.Printf("connect to db succ \n")
	//QueryRow(Db)
	//QuerySelect(Db)
	//Query(Db)
    Insert(Db)
}