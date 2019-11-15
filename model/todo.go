package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Todo struct {
	ID     int    `json:"id" gorm:"primary_key auto_increment"`
	Name   string `json:"name"`
	Done   bool   `json:"done" gorm:"default: false"`
	UserID uint   `json:"user_id"`
}

func CreateTodo(todo *Todo, DB *gorm.DB) error {
	DB.Create(todo)

	return nil
}

func GetTodos(t *Todo, DB *gorm.DB) ([]Todo, error) {
	var todos []Todo

	DB.Where(t).Find(&todos)
	return todos, nil
}

func UpdateTodo(todo *Todo, DB *gorm.DB) error {
	if rows := DB.Model(todo).Update(map[string]interface{}{
		"name": todo.Name,
		"done": todo.Done,
	}).RowsAffected; rows == 0 {
		return fmt.Errorf("could not find %v", todo)
	}

	return nil
}

func DeleteTodo(todo *Todo, DB *gorm.DB) error {
	if rows := DB.Where(todo).Delete(&Todo{}).RowsAffected; rows == 0 {
		return fmt.Errorf("could not find %v", todo)
	}

	return nil
}
