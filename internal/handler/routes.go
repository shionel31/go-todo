package handler

import (
	"github.com/labstack/echo/v4"
)

func NewTodoRouter(todoHH TodoHandler) *echo.Echo {
	e := echo.New()
	todos := e.Group("/todos")

	todos.GET("", todoHH.List)
	todos.GET("/:todoID", todoHH.Get)
	todos.POST("", todoHH.Create)
	todos.PUT("/:todoID", todoHH.Update)
	todos.DELETE("/:todoID", todoHH.Delete)

	return e
}
