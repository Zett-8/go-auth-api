package main

import "go-auth-api/model"

func main() {
	model.Init()
	defer model.DB.Close()

	router := newRouter()
	router.Logger.Fatal(router.Start(":8080"))
}
