package main

import (
	"fmt"
	"go-todo/internal/domain/model"
	"go-todo/internal/infrastructure/database"
)

func main() {
	dbConn := database.NewDB()
	defer fmt.Println(("マイグレーションが正常に実行されました"))
	defer database.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.Todo{})
}
