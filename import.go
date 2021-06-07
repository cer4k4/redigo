package main

import (
        "log"
	 "database/sql"
	 _ "github.com/go-sql-driver/mysql"
)


func main() {
	db, err := sql.Open("mysql", "golang:golang123@(localhost)/golang_test")
	if err != nil{
		panic(err.Error())
	}
	defer db.Close()
	log.Println(db)
}
