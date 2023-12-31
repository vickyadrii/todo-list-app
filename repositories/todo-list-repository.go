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

func GetAllTodoLists(newList *[]models.TodoList) error {
	result := config.DB.Find(&newList)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetTodoList(list *models.TodoList) error {
	result := config.DB.Take(&list)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateTodoList(list *models.TodoList) error {
	result := config.DB.Save(&list)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteTodoList(list *models.TodoList, id string) error {
	err := config.DB.First(&list, id).Error
	if err != nil {
		return err
	}
	err = config.DB.Delete(&list, id).Error

	if err != nil {

		return err
	}
	return nil
}
