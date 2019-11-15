package model

type User struct {
	ID       int    `json:"id" gorm:"primary_key auto_increment"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Todos    []Todo
}

func CreateUser(user *User) error {
	DB, err := connectDB()
	if err != nil {
		return err
	}
	defer DB.Close()

	DB.Create(user)

	return nil
}

func GetUser(u *User) (User, error) {
	var user User
	var todos []Todo

	DB, err := connectDB()
	if err != nil {
		return User{}, err
	}
	defer DB.Close()

	DB.Where(u).First(&user)

	DB.Where(Todo{UserID: uint(user.ID)}).Find(&todos)
	user.Todos = todos

	return user, nil
}
