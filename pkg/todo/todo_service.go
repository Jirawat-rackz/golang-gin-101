package todo

import "github.com/jirawat-rackz/golang-gin-101/pkg/model"

type ITodoService interface {
	GetAllTodo() ([]model.Todo, error)
	InsertTodo(todo model.Todo) (model.Todo, error)
}

type TodoService struct {
	TodoRepository ITodoRepository
}

func NewTodoService() *TodoService {
	return &TodoService{
		TodoRepository: NewTodoRepository(),
	}
}

func (service *TodoService) GetAllTodo() ([]model.Todo, error) {
	result, err := service.TodoRepository.GetAllTodo()
	return result, err
}

func (service *TodoService) InsertTodo(todo model.Todo) (model.Todo, error) {
	result, err := service.TodoRepository.InsertTodo(todo)
	return result, err
}
