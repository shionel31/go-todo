package handler

import (
	"go-todo/internal/application/usecase"
	"go-todo/internal/domain/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TodoHandler interface {
	List(c echo.Context) error
	Get(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

// パッケージプライベートな構造体のため、小文字始まり
type todoHandler struct {
	todoUC usecase.TodoUsecase
}

// インスタンス化用のファクトリメソッド
func NewTodoHandler(uc usecase.TodoUsecase) TodoHandler {
	return &todoHandler{
		todoUC: uc,
	}
}

func (th *todoHandler) List(c echo.Context) error {
	todos, err := th.todoUC.ListTodos()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, todos)
}

func (th *todoHandler) Get(c echo.Context) error {
	ID := c.Param("todoID")
	todoID, _ := strconv.Atoi(ID)
	todo, err := th.todoUC.GetTodo(uint(todoID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, todo)
}

func (th *todoHandler) Create(c echo.Context) error {
	todo := model.Todo{}
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := th.todoUC.CreateTodo(todo); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, todo)
}

func (th *todoHandler) Update(c echo.Context) error {
	ID := c.Param("todoID")
	todoID, _ := strconv.Atoi(ID)

	// Bind用todoを生成
	todo := model.Todo{}
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := th.todoUC.UpdateTodo(uint(todoID), todo); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, todo)
}

func (th *todoHandler) Delete(c echo.Context) error {
	ID := c.Param("todoID")
	todoID, _ := strconv.Atoi(ID)

	if err := th.todoUC.DeleteTodo(uint(todoID)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
