package repositoryIF

import "go-todo/internal/domain/model"

type TodoRepositoryIF interface {
	FetchAll(todos *[]model.Todo) error
	FindByID(todoID uint, todo *model.Todo) error
	InsertOne(todo *model.Todo) error
	UpdateOne(todoID uint, todo *model.Todo) error
	DeleteOne(todoID uint) error
}
