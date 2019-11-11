package model

type User struct {
	ID       int    `json:"id" gorm:"primary_key auto_increment"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Todos    []Todo
}

func CreateUser(user *User) {
	db.Create(user)
}

func GetUser(u *User) User {
	var user User
	var todos []Todo
	db.Where(u).First(&user)

	db.Where(Todo{UserID: uint(user.ID)}).Find(&todos)
	user.Todos = todos

	return user
}
