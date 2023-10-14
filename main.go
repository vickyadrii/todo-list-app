package main

import (
	"fmt"
	"todo-list-app/config"
	"todo-list-app/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// config.LoadEnv()
	config.InitDB()
	e := echo.New()
	routes.TodoListRoutes(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", 8000)))
}
