package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DBConfig = "host=db port=5432 user=postgres dbname=app password=postgres sslmode=disable"

func Init() {
	DB, err := gorm.Open("postgres", DBConfig)

	if err != nil {
		panic(err.Error())
	}

	defer DB.Close()

	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Todo{})
}

func connectDB() (*gorm.DB, error) {
	DB, err := gorm.Open("postgres", DBConfig)
	return DB, err
}
