package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"syscall"

	"taskqueue/internal/application/handlers"
	"taskqueue/internal/domain/entities"

	"golang.org/x/term"
)

type CLI struct {
	userApp     *handlers.UserApp
	authApp     *handlers.AuthApp
	currentUser *entities.User
}

func NewCLI(userApp *handlers.UserApp, authApp *handlers.AuthApp) *CLI {
	return &CLI{
		userApp: userApp,
		authApp: authApp,
	}
}

func (cli *CLI) Start() {
	fmt.Println("=== TaskQueue System ===")

	for {
		if cli.currentUser == nil {
			cli.loginPrompt()
		} else {
			cli.mainMenu()
		}
	}
}

func (cli *CLI) loginPrompt() {
	fmt.Print("Email: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	email := strings.TrimSpace(scanner.Text())

	fmt.Print("Password: ")
	passwordBytes, _ := term.ReadPassword(int(syscall.Stdin))
	password := string(passwordBytes)
	fmt.Println()

	user, token, err := cli.authApp.Login(context.Background(), email, password)
	if err != nil {
		fmt.Printf("Ошибка входа: %v\n", err)
		return
	}

	cli.currentUser = user
	fmt.Printf("Добро пожаловать, %s %s! Роль: %s\n",
		user.Username, user.Surname, user.Role)

	os.Setenv("AUTH_TOKEN", token)
}

func (cli *CLI) mainMenu() {
	fmt.Printf("\n=== Главное меню (Роль: %s) ===\n", cli.currentUser.Role)

	switch cli.currentUser.Role {
	case "super":
		cli.superUserMenu()
	case "medium":
		cli.mediumUserMenu()
	case "regular":
		cli.regularUserMenu()
	}
}

func (cli *CLI) superUserMenu() {
	fmt.Println("1. Создать пользователя")
	fmt.Println("2. Просмотреть всех пользователей")
	fmt.Println("3. Создать задачу")
	fmt.Println("4. Просмотреть задачи")
	fmt.Println("0. Выйти")

	fmt.Print("Выберите действие: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choice := strings.TrimSpace(scanner.Text())

	switch choice {
	case "1":
		cli.createUserPrompt()
	case "2":
		cli.listUsers()
	case "3":
		fmt.Println("Создание задач - в разработке")
	case "4":
		fmt.Println("Просмотр задач - в разработке")
	case "0":
		cli.logout()
	default:
		fmt.Println("Неверный выбор")
	}
}

func (cli *CLI) createUserPrompt() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Имя: ")
	scanner.Scan()
	username := strings.TrimSpace(scanner.Text())

	fmt.Print("Фамилия: ")
	scanner.Scan()
	surname := strings.TrimSpace(scanner.Text())

	fmt.Print("Отчество: ")
	scanner.Scan()
	patronymic := strings.TrimSpace(scanner.Text())

	fmt.Print("Email: ")
	scanner.Scan()
	email := strings.TrimSpace(scanner.Text())

	fmt.Print("Пароль: ")
	passwordBytes, _ := term.ReadPassword(int(syscall.Stdin))
	password := string(passwordBytes)
	fmt.Println()

	fmt.Print("Роль (medium/regular): ")
	scanner.Scan()
	role := strings.TrimSpace(scanner.Text())

	userID, err := cli.userApp.CreateUserBySuper(
		context.Background(),
		username, surname, patronymic, email, password, role,
	)

	if err != nil {
		fmt.Printf("Ошибка создания пользователя: %v\n", err)
		return
	}

	fmt.Printf("Пользователь успешно создан с ID: %d\n", userID)
}

func (cli *CLI) listUsers() {
	fmt.Println("=== Список пользователей ===")
	fmt.Println("Функция в разработке - нужно добавить метод GetAllUsers в UserApp")
	fmt.Println("Для полной реализации потребуется:")
	fmt.Println("1. Добавить метод GetAll в UserRepository")
	fmt.Println("2. Добавить GetAllUsers в UserApp")
	fmt.Println("3. Реализовать отображение списка здесь")
}

func (cli *CLI) mediumUserMenu() {
	fmt.Println("1. Просмотреть задачи от супер пользователя")
	fmt.Println("2. Создать задачу для обычного пользователя")
	fmt.Println("0. Выйти")

	fmt.Print("Выберите действие: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choice := strings.TrimSpace(scanner.Text())

	switch choice {
	case "0":
		cli.logout()
	default:
		fmt.Println("Функционал в разработке")
	}
}

func (cli *CLI) regularUserMenu() {
	fmt.Println("1. Просмотреть мои задачи")
	fmt.Println("2. Отправить отчет")
	fmt.Println("0. Выйти")

	fmt.Print("Выберите действие: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choice := strings.TrimSpace(scanner.Text())

	switch choice {
	case "0":
		cli.logout()
	default:
		fmt.Println("Функционал в разработке")
	}
}

func (cli *CLI) logout() {
	cli.currentUser = nil
	os.Unsetenv("AUTH_TOKEN")
	fmt.Println("Вы вышли из системы")
}
