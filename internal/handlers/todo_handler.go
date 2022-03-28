package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"todo-api/internal/request"
	"todo-api/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type todoHandler struct {
	todoService services.TodoService
}

func NewTodoHandler(todoService services.TodoService) *todoHandler {
	return &todoHandler{todoService}
}

func (h todoHandler) GetAllTodoHandler(c *gin.Context) {
	todos, _ := h.todoService.GetAllTodo()

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   todos,
	})
}

func (h todoHandler) FindByIDTodoHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, _ := h.todoService.FindByIDTodo(id)
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   todo,
	})
}

func (h todoHandler) StoreTodoHandler(c *gin.Context) {
	var todoRequest request.TodoRequest

	err := c.ShouldBindJSON(&todoRequest)

	fmt.Println(err)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on filled %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	todo, err := h.todoService.CreateTodo(todoRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": todo,
	})
}

func (h todoHandler) UpdateTodoHandler(c *gin.Context) {
	var todoRequest request.TodoRequest

	err := c.ShouldBindJSON(&todoRequest)

	id, _ := strconv.Atoi(c.Param("id"))

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on filled %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	todo, err := h.todoService.UpdateTodo(id, todoRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": todo,
	})
}

func (h todoHandler) DeleteTodoHandler(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := h.todoService.DeleteTodo(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": todo,
	})
}
