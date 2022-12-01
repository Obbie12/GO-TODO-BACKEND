package models

import (
	"fmt"
	"go-todo-backend/database"
)

func GetAllTodos(todo *[]Todo) (err error) {
	if err = database.Database.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func CreateATodo(todo *Todo) (err error) {
	if err = database.Database.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func GetATodo(todo *Todo, id int) (err error) {
	if err := database.Database.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func UpdateATodo(todo *Todo, id int) (err error) {
	fmt.Println(todo)
	database.Database.Save(todo)
	return nil
}

func DeleteATodo(todo *Todo, id int) (err error) {
	database.Database.Where("id = ?", id).Delete(todo)
	return nil
}
