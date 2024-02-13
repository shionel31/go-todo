package main

import (
	"go-todo/internal/application/usecase"
	"go-todo/internal/handler"
	"go-todo/internal/infrastructure/database"
	"go-todo/internal/infrastructure/repository"
)

func main() {
	db := database.NewDB()
	todoRepository := repository.NewTodoRepository(db)
	todoUsecase := usecase.NewTodoUsecase(todoRepository)
	todoHandler := handler.NewTodoHandler(*todoUsecase)
	e := handler.NewTodoRouter(todoHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
