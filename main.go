package main

import "go-auth-api/model"

func main() {
	model.Init()

	router := newRouter()
	router.Logger.Fatal(router.Start(":8080"))
}
