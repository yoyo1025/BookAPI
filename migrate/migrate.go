package main

import (
	"BookAPI/db"
	"BookAPI/model"
	"fmt"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Book{}) // データベースに反映させたいモデル構造を渡す
}