package main

import (
	"os"
	"todo-list-app/config"
	"todo-list-app/routes"

	"github.com/labstack/echo/v4"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		// Default to port 8080 if PORT is not set
		port = "8080"
	}
	return port
}

func main() {
	config.LoadEnv()
	config.InitDB()
	e := echo.New()
	routes.TodoListRoutes(e)
	e.Logger.Fatal(e.Start(":" + getPort())) 
	// e.Start(":45669")
}
