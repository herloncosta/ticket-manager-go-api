package main

import (
	"github.com/herloncosta/ticket-manager/internal/config"
	"github.com/herloncosta/ticket-manager/internal/db"
	"github.com/herloncosta/ticket-manager/internal/router"
	"log"
	"os"
)

func main() {
	cfg := config.Load()
	database, err := db.NewSQLite(cfg.DBPath)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer database.Close()

	db.RunMigrations(database)

	r := router.Setup(database)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on port %s", port)
	r.Run(":" + port)
}
