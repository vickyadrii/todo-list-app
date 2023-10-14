package routes

import (
	"todo-list-app/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoutes(e *echo.Echo) {
	e.Use(middleware.Logger())
	eAuth := e.Group("")
	eAuth.POST("/lists", controllers.CreateTodoList)
	eAuth.GET("/lists", controllers.GetAllTodoList)
}
