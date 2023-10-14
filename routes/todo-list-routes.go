package routes

import (
	"todo-list-app/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoutes(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	eAuth := e.Group("")
	eAuth.POST("/lists", controllers.CreateTodoListController)
	eAuth.GET("/lists", controllers.GetAllTodoListController)
	eAuth.GET("lists/:id", controllers.GetTodoListController)
	eAuth.PUT("/lists/:id", controllers.UpdateTodoListController)
}
