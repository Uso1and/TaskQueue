package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"taskqueue/internal/application/handlers"
	"taskqueue/internal/domain/entities"
	"taskqueue/internal/domain/repositories"
	"taskqueue/internal/infrastructure/persistence/postgresql"
	"taskqueue/internal/interfaces/http/routes"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	var cliMode = flag.Bool("cli", false, "Запуск в режиме CLI")
	var createSuper = flag.Bool("create-super", false, "Создать супер пользователя")
	flag.Parse()

	if err := postgresql.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	if err := postgresql.RunMigrations(postgresql.DB); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	userRepo := postgresql.NewUserRepository(postgresql.DB)
	userApp := handlers.NewUserApp(userRepo)

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "default-secret-key"
	}
	authApp := handlers.NewAuthApp(userRepo, jwtSecret)

	if *createSuper {
		createSuperUser(userRepo)
		return
	}

	if *cliMode {
		cli := NewCLI(userApp, authApp)
		cli.Start()
	} else {

		router := routes.SetupRouter()
		log.Println("Server starting on :8080")
		if err := router.Run(":8080"); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}
}

func createSuperUser(userRepo repositories.UserRepository) {
	fmt.Print("Email супер пользователя: ")
	var email string
	fmt.Scanln(&email)

	fmt.Print("Пароль: ")
	var password string
	fmt.Scanln(&password)

	fmt.Print("Имя: ")
	var username string
	fmt.Scanln(&username)

	fmt.Print("Фамилия: ")
	var surname string
	fmt.Scanln(&surname)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Ошибка хеширования пароля: %v", err)
	}

	user := &entities.User{
		Username:   username,
		Surname:    surname,
		Patronymic: "",
		Email:      email,
		Password:   string(hashedPassword),
		Role:       "super",
	}

	err = userRepo.Create(context.Background(), user)
	if err != nil {
		log.Fatalf("Ошибка создания пользователя: %v", err)
	}

	fmt.Printf("Супер пользователь успешно создан с ID: %d\n", user.ID)
}
