package main

import (
	"go-auth-api/handler"
	"go-auth-api/model"
)

func main() {
	model.Init()

	router := handler.NewRouter()
	router.Logger.Fatal(router.Start(":8080"))
}
