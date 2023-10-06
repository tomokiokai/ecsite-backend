package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/model"
)

func main() {
	dbConn := db.NewDB()
	defer db.CloseDB(dbConn)

	// 既存のモデルと新しい Favorite モデルをマイグレートします
	err := dbConn.AutoMigrate(&model.User{}, &model.Task{}, &model.Blog{}, &model.Shop{}, &model.Favorite{})
	if err != nil {
		fmt.Println("Migration failed:", err)
		return
	}

	fmt.Println("Successfully Migrated")
}

