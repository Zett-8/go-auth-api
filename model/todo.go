package model

import "fmt"

type Todo struct {
	ID     int    `json:"id" gorm:"primary_key auto_increment"`
	Name   string `json:"name"`
	Done   bool   `json:"done" gorm:"default: false"`
	UserID uint   `json:"user_id"`
}

func CreateTodo(todo *Todo) error {
	DB, err := connectDB()
	if err != nil {
		return err
	}
	defer DB.Close()

	DB.Create(todo)

	return nil
}

func GetTodos(t *Todo) ([]Todo, error) {
	var todos []Todo

	DB, err := connectDB()
	if err != nil {
		return []Todo{}, err
	}
	defer DB.Close()

	DB.Where(t).Find(&todos)
	return todos, nil
}

func UpdateTodo(todo *Todo) error {
	DB, err := connectDB()
	if err != nil {
		return err
	}
	defer DB.Close()

	if rows := DB.Model(todo).Update(map[string]interface{}{
		"name": todo.Name,
		"done": todo.Done,
	}).RowsAffected; rows == 0 {
		return fmt.Errorf("could not find %v", todo)
	}

	return nil
}

func DeleteTodo(todo *Todo) error {
	DB, err := connectDB()
	if err != nil {
		return err
	}
	defer DB.Close()

	if rows := DB.Where(todo).Delete(&Todo{}).RowsAffected; rows == 0 {
		return fmt.Errorf("could not find %v", todo)
	}

	return nil
}
