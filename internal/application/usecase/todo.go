package usecase

import (
	"go-todo/internal/domain/model"
	"go-todo/internal/domain/repositoryIF"
)

// リポジトリインターフェースはドメイン層から取ってくる
type TodoUsecase struct {
	todoRepoIF repositoryIF.TodoRepositoryIF
}

// type TodoUsecase interface {
// 	ListTodos() ([]model.Todo, error)
// 	GetTodoByID(todoID uint, todo *model.Todo) error
// 	CreateTodo(todo *model.Todo) error
// 	UpdateTodo(todoID uint, todo *model.Todo) error
// 	DeleteTodo(todoID uint) error
// }

// サーバー側で呼び出されるインスタンス化用のファクトリメソッド
func NewTodoUsecase(repoIF repositoryIF.TodoRepositoryIF) *TodoUsecase {
	return &TodoUsecase{
		todoRepoIF: repoIF,
	}
}

func (uc *TodoUsecase) ListTodos() ([]model.Todo, error) {
	todos := []model.Todo{}
	if err := uc.todoRepoIF.FetchAll(&todos); err != nil {
		return nil, err
	}
	return todos, nil
}

func (uc *TodoUsecase) GetTodo(todoID uint) (model.Todo, error) {
	todo := model.Todo{}
	if err := uc.todoRepoIF.FindByID(todoID, &todo); err != nil {
		return model.Todo{}, err
	}
	return todo, nil
}

func (uc *TodoUsecase) CreateTodo(todo model.Todo) error {
	if err := uc.todoRepoIF.InsertOne(&todo); err != nil {
		return err
	}
	return nil
}

func (uc *TodoUsecase) UpdateTodo(todoID uint, todo model.Todo) error {
	if err := uc.todoRepoIF.UpdateOne(todoID, &todo); err != nil {
		return err
	}
	return nil
}

func (uc *TodoUsecase) DeleteTodo(todoID uint) error {
	if err := uc.todoRepoIF.DeleteOne(todoID); err != nil {
		return err
	}
	return nil
}
