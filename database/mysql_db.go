package database

import (
	"fmt"
	"log"

	"todo-api/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysqlDB() (*gorm.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"
	result, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB CONNECTION ERROR")
	}
	fmt.Println("Database Connection Succeed")
	result.AutoMigrate(&entity.Todo{})
	return result, err
}
