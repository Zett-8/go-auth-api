package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	ID       int    `json:"id" gorm:"primary_key auto_increment"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Todos    []Todo
}

func CreateUser(user *User, DB *gorm.DB) error {
	DB.Create(user)

	return nil
}

func GetUser(u *User, DB *gorm.DB) (User, error) {
	var user User
	var todos []Todo

	DB.Where(u).First(&user)

	DB.Where(Todo{UserID: uint(user.ID)}).Find(&todos)
	user.Todos = todos

	return user, nil
}
