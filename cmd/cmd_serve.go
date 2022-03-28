package main

import (
	"github.com/gin-gonic/gin"

	"todo-api/internal/database"
	"todo-api/internal/routes"
)

func main() {
	db, _ := database.InitMysqlDB()

	router := gin.Default()
	routes.TodoRoute(router, db) //Added all auth routes
	router.Run()
}
