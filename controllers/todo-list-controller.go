package controllers

import (
	"net/http"
	"todo-list-app/models"
	"todo-list-app/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateTodoListController(c echo.Context) error {
	var todoRequest models.TodoList
	c.Bind(&todoRequest)

	err := repositories.AddTodoList(&todoRequest)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response{
			Message: err.Error(),
			Status:  http.StatusNotFound,
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "Add data list successfully!",
		Status:  http.StatusCreated,
		Data:    todoRequest,
	})
}

func GetAllTodoListController(c echo.Context) error {
	var lists []models.TodoList

	err := repositories.GetAllTodoLists(&lists)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
			Data:    nil,
		})

	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "Get all data list successfully!",
		Status:  http.StatusOK,
		Data:    lists,
	})
}

func GetTodoListController(c echo.Context) error {
	list := models.TodoList{}
	c.Bind(&list)

	err := repositories.GetTodoList(&list)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
			Data:    nil,
		})

	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "Get data list successfully!",
		Status:  http.StatusOK,
		Data:    list,
	})
}

func UpdateTodoListController(c echo.Context) error {
	list := models.TodoList{}
	c.Bind(&list)

	err := repositories.UpdateTodoList(&list)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "Update data successfully!",
		Status:  http.StatusOK,
		Data:    list,
	})
}

func DeleteTodoListController(c echo.Context) error {
	id := c.Param("id")

	err := repositories.DeleteTodoList(&models.TodoList{}, id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, models.Response{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Data:    nil,
			})
		}
		return c.JSON(http.StatusInternalServerError, models.Response{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "Delete data successfully!",
		Status:  http.StatusOK,
	})
}
