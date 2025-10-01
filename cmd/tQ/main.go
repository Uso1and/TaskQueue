package main

import (
	"log"
	"taskqueue/internal/infrastructure/persistence/postgresql"
	"taskqueue/internal/interfaces/http/routes"
)

func main() {

	if err := postgresql.InitDB(); err != nil {
		log.Printf("Failed to initialize database: %v", err)
	}

	if err := postgresql.RunMigrations(postgresql.DB); err != nil {
		log.Fatal(err)
	}

	r := routes.SetupRouter()

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
