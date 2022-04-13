package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"todo-api/internal/database"
	"todo-api/internal/handlers"
	"todo-api/internal/repository"
	"todo-api/internal/services"
)

func main() {
	db, _ := database.InitMysqlDB()
	router := gin.Default()

	setupTodo(router, db)

	router.Run()
}

func setupTodo(router gin.IRouter, db *gorm.DB) {
	repo := repository.NewTodoRepository(db)
	svc := services.NewTodoService(repo)

	handler := handlers.NewTodoHandler(svc)

	router.GET("/blabla", handler.GetAllTodoHandler)
}
