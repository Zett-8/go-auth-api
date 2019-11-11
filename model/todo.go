package model

type Todo struct {
	ID     int    `json:"id" gorm:"primary_key auto_increment"`
	Name   string `json:"name"`
	Done   bool   `json:"done" gorm:"default: false"`
	UserID uint   `json:"user_id"`
}

func CreateTodo(todo *Todo) {
	db.Create(todo)
}

func GetTodo(todo Todo) {
	db.Find(todo)
}
