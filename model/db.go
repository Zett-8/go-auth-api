package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func Init() {
	var err error

	DB, err = gorm.Open("postgres", "host=db port=5432 user=postgres dbname=app password=postgres sslmode=disable")

	if err != nil {
		panic(err.Error())
	}

	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Todo{})
}
