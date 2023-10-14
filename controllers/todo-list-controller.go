package controllers

import (
	"net/http"
	"todo-list-app/models"
	"todo-list-app/repositories"

	"github.com/labstack/echo/v4"
)

func CreateTodoList(c echo.Context) error {
	var todoRequest models.TodoList
	c.Bind(&todoRequest)

	err := repositories.AddTodoList(&todoRequest)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "Add data list successfully!",
		Status:  http.StatusCreated,
		Data:    todoRequest,
	})
}

func GetAllTodoList(c echo.Context) error {
	var lists []models.TodoList

	err := repositories.GetTodoList(&lists)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
			Data:    nil,
		})

	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "Get Data List Successfully!",
		Status:  http.StatusOK,
		Data:    lists,
	})
}
