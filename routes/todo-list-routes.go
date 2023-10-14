package routes

import (
	"todo-list-app/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func TodoListRoutes(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.POST("/lists", controllers.CreateTodoListController)
	e.GET("/lists", controllers.GetAllTodoListController)
	e.GET("lists/:id", controllers.GetTodoListController)
	e.PUT("/lists/:id", controllers.UpdateTodoListController)
	e.DELETE("/lists/:id", controllers.DeleteTodoListController)
}
