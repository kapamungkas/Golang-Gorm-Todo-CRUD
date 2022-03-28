package main

import (
	"todo-api/database"
	"todo-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db, _ := database.InitMysqlDB()

	router := gin.Default()
	routes.TodoRoute(router, db) //Added all auth routes
	router.Run()
}
