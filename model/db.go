package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func App_db() *gorm.DB {
	var err error

	db, err = gorm.Open("postgres", "host=db port=5432 user=postgres dbname=app password=postgres sslmode=disable")
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Todo{})

	return db
}
