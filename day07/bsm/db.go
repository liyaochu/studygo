package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//跟数据库相关
var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:root@tcp(127.0.0.1)/go_test"
	db, err = sqlx.Connect("mysql", dsn)
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

//查数据
func queryAllBook() (bookList []*Book, err error) {

	sqlStr := "select id ,title,price from book"
	err = db.Select(&bookList, sqlStr)
	if err != nil {
		fmt.Println("查询所有书籍失败")
		return
	}
	return
}

//查数据
func insertBook(title string ,price float64) ( err error) {

	sqlStr := "insert into  book(title,price) values(?,?)"
	_, err = db.Exec(sqlStr, title, price)
	if err != nil {
		fmt.Println("查询所有书籍失败")
		return
	}
	return
}
