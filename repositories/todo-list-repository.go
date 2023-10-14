package repositories

import (
	"todo-list-app/config"
	"todo-list-app/models"
)

func AddTodoList(todoList *models.TodoList) error {
	result := config.DB.Create(&todoList)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetTodoList(newList *[]models.TodoList) error {
	result := config.DB.Find(&newList)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
