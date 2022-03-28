package repository

import (
	"todo-api/internal/entity"

	"gorm.io/gorm"
)

type TodoRepository interface {
	GetAllTodo() ([]entity.Todo, error)
	FindByIDTodo(id int) (entity.Todo, error)
	CreateTodo(todoEntity entity.Todo) (entity.Todo, error)
	UpdateTodo(id int, todoEntity entity.Todo) (entity.Todo, error)
	DeleteTodo(id int) (entity.Todo, error)
}

type todoRepository struct {
	databBase *gorm.DB
}

func NewTodoRepository(databBase *gorm.DB) *todoRepository {
	return &todoRepository{databBase}
}

func (t todoRepository) GetAllTodo() ([]entity.Todo, error) {
	var todos []entity.Todo

	err := t.databBase.Find(&todos).Error

	return todos, err
}

func (t todoRepository) FindByIDTodo(id int) (entity.Todo, error) {
	var todos entity.Todo

	err := t.databBase.Find(&todos, id).Error

	return todos, err

}

func (t todoRepository) CreateTodo(entityTodo entity.Todo) (entity.Todo, error) {
	err := t.databBase.Create(&entityTodo).Error
	return entityTodo, err
}

func (t todoRepository) UpdateTodo(id int, entityTodo entity.Todo) (entity.Todo, error) {
	var todos entity.Todo

	err := t.databBase.Model(&todos).Where("id = ?", id).Updates(entityTodo).Error

	return entityTodo, err
}

func (t todoRepository) DeleteTodo(id int) (entity.Todo, error) {
	var todos entity.Todo

	err := t.databBase.Where("id = ?", id).Delete(&todos).Error

	return todos, err
}
