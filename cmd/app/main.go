package main

import (
	"log"

	_ "online-library/docs" // swagger

	"online-library/internal/router"
	"online-library/pkg/config"
	"online-library/pkg/db"
)

// @title           Online Library API
// @version         1.0
// @description     This is a Online Library API with CRUD operations.
// @host            localhost:8080
// @BasePath        /
func main() {
	cfg := config.LoadConfig()
	dbConn, err := db.InitDB(*cfg)
	if err != nil {
		log.Fatalf("Database Init Failed : %v", err)
	}
	defer dbConn.Close()

	log.Println("APP Started : ", cfg.AppName)

	r := router.SetupRouter(dbConn)
	r.Run(":8080")
}
