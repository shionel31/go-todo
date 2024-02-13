package repository

import (
	"fmt"
	"go-todo/internal/domain/model"
	"go-todo/internal/domain/repositoryIF"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// パッケージプライベートな構造体のため、小文字始まり
type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) repositoryIF.TodoRepositoryIF {
	return &todoRepository{db}
}

func (todoRepo *todoRepository) FetchAll(todos *[]model.Todo) error {
	if err := todoRepo.db.Find(&todos).Error; err != nil {
		return err
	}
	return nil
}

func (todoRepo *todoRepository) FindByID(todoID uint, todo *model.Todo) error {
	if err := todoRepo.db.First(todo, todoID).Error; err != nil {
		return err
	}
	return nil
}

func (todoRepo *todoRepository) InsertOne(todo *model.Todo) error {
	if err := todoRepo.db.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func (todoRepo *todoRepository) UpdateOne(todoID uint, todo *model.Todo) error {
	result := todoRepo.db.Model(todo).Clauses(clause.Returning{}).Where("id=?", todoID).Updates(
		map[string]interface{}{
			"title":  todo.Title,
			"status": todo.Status,
		},
	)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (todoRepo *todoRepository) DeleteOne(todoID uint) error {
	result := todoRepo.db.Where("id=?", todoID).Delete(&model.Todo{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
