package main

import "go-auth-api/model"

func main() {
	db := model.App_db()
	defer db.Close()

	router := newRouter()
	router.Logger.Fatal(router.Start(":8080"))
}
