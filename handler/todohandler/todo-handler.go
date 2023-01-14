package todohandler

import (
	"github.com/gin-gonic/gin"
	"github.com/jirawat-rackz/golang-gin-101/pkg/model"
	"github.com/jirawat-rackz/golang-gin-101/pkg/todo"
)

type TodoHandler struct {
	TodoService todo.ITodoService
}

func NewTodoHandler() *TodoHandler {
	return &TodoHandler{
		TodoService: todo.NewTodoService(),
	}
}

func (handler *TodoHandler) GetAllTodo(c *gin.Context) {
	result, err := handler.TodoService.GetAllTodo()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(200, result)
}

func (handler *TodoHandler) PostTodo(c *gin.Context) {
	var todo model.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(500, err)
	}

	result, err := handler.TodoService.InsertTodo(todo)
	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(201, result)
}
