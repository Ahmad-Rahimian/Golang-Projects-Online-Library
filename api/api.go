package api

import (
	"log"

	"online-library/internal/router"
	pkgCF "online-library/pkg/config"
	pkgDB "online-library/pkg/db"
)

func Start() {
	// load config
	cfg := pkgCF.LoadConfig()

	// connect to db
	db, err := pkgDB.InitDB(*cfg)
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// setup router
	r := router.SetupRouter(db, cfg, cfg.DB.JWTSecret)

	// run server
	r.Run(":8080")
}
