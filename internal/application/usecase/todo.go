package usecase

import "github.com/shionel31/go-todo/internal/interface"

type TodoUsecase struct {
	todoRepository interface.TodoRepositoryIF
}
