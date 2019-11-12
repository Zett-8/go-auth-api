package model

import "fmt"

type Todo struct {
	ID     int    `json:"id" gorm:"primary_key auto_increment"`
	Name   string `json:"name"`
	Done   bool   `json:"done" gorm:"default: false"`
	UserID uint   `json:"user_id"`
}

func CreateTodo(todo *Todo) {
	db.Create(todo)
}

func GetTodo(t *Todo) []Todo {
	var todos []Todo
	db.Where(t).Find(&todos)
	return todos
}

func DeleteTodo(todo *Todo) error {
	if rows := db.Where(todo).Delete(&Todo{}).RowsAffected; rows == 0 {
		return fmt.Errorf("could not find %v", todo)
	}

	return nil
}
