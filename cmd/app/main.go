package main

import (
	"online-library/api"
	_ "online-library/docs"
)

// @title           Online Library API
// @version         1.0
// @description     This is a Online Library API with CRUD operations.
// @host            localhost:8080
// @BasePath        /
func main() {
	api.Start()
}
