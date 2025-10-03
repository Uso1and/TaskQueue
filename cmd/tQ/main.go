package main

import (
	"log"
	"taskqueue/internal/infrastructure/persistence/postgresql"
	"taskqueue/internal/interfaces/http/routes"
)

func main() {

	if err := postgresql.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	if err := postgresql.RunMigrations(postgresql.DB); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	router := routes.SetupRouter()

	log.Println("Server starting on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
