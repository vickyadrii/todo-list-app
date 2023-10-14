package config

import (
	"fmt"
	"os"
	"todo-list-app/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func InitDB() {
	connectDB := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DBUSER"),
		os.Getenv("DBPASSWORD"),
		os.Getenv("DBHOST"),
		os.Getenv("DBPORT"),
		os.Getenv("DBNAME"),
	)
	DB, err = gorm.Open(mysql.Open(connectDB), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to database")
	}
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.TodoList{})
}
