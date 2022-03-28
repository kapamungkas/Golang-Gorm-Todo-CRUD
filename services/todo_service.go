package services

import (
	"todo-api/entity"
	"todo-api/repository"
	"todo-api/request"
)

type TodoService interface {
	GetAllTodo() ([]entity.Todo, error)
	FindByIDTodo(id int) (entity.Todo, error)
	CreateTodo(todoRequest request.TodoRequest) (entity.Todo, error)
	UpdateTodo(id int, todoRequest request.TodoRequest) (entity.Todo, error)
	DeleteTodo(id int) (entity.Todo, error)
	// FindByIDTodo() (string, error)
	// FindByIDTodo() (Todo, error)
}

type todoService struct {
	todoRepository repository.TodoRepository
}

func NewTodoService(todoRepository repository.TodoRepository) *todoService {
	return &todoService{todoRepository}
}

func (t todoService) GetAllTodo() ([]entity.Todo, error) {
	// var todo []Todo
	todo, err := t.todoRepository.GetAllTodo()
	return todo, err
}

func (t todoService) FindByIDTodo(id int) (entity.Todo, error) {
	todo, err := t.todoRepository.FindByIDTodo(id)
	return todo, err
}

func (t todoService) CreateTodo(todoRequest request.TodoRequest) (entity.Todo, error) {
	todo := entity.Todo{
		Todo:        todoRequest.Todo,
		Description: todoRequest.Description,
	}

	res, err := t.todoRepository.CreateTodo(todo)
	return res, err
}

func (t todoService) UpdateTodo(id int, todoRequest request.TodoRequest) (entity.Todo, error) {
	todo := entity.Todo{
		Todo:        todoRequest.Todo,
		Description: todoRequest.Description,
	}

	res, err := t.todoRepository.UpdateTodo(id, todo)
	return res, err
}

func (t todoService) DeleteTodo(id int) (entity.Todo, error) {

	res, err := t.todoRepository.DeleteTodo(id)
	return res, err
}
