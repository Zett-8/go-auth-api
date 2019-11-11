package model

type User struct {
	ID       int    `json:"id" gorm:"primary_key auto_increment"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func CreateUser(user User) {
	db.Create(user)
}

func GetUser(id int) User {
	var user User
	db.Where(User{ID: id}).First(&user)
	return user
}
