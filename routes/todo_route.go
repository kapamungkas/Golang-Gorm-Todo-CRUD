package routes

import (
	"todo-api/handlers"
	"todo-api/services"

	"todo-api/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TodoRoute(route *gin.Engine, db *gorm.DB) {
	todo := route.Group("/todo")
	todo_repository := repository.NewTodoRepository(db)
	todo_service := services.NewTodoService(todo_repository)
	h := handlers.NewTodoHandler(todo_service)
	todo.GET("/", h.GetAllTodoHandler)
	todo.GET("/:id", h.FindByIDTodoHandler)
	todo.POST("/", h.StoreTodoHandler)
	todo.PUT("/:id", h.UpdateTodoHandler)
	todo.DELETE("/:id", h.DeleteTodoHandler)
}
