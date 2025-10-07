package commands

import (
	"context"
	"errors"
	"strings"

	"taskqueue/internal/domain/entities"
	"taskqueue/internal/domain/repositories"

	"golang.org/x/crypto/bcrypt"
)

type CreateUser struct {
	Username   string
	Surname    string
	Patronymic string
	Email      string
	Password   string
	Role       string
}

type CreateUserHandler struct {
	Users repositories.UserRepository
}

func NewCreateUserHandler(users repositories.UserRepository) *CreateUserHandler {
	return &CreateUserHandler{Users: users}
}

func (h *CreateUserHandler) Handle(ctx context.Context, cmd CreateUser) (int, error) {
	if cmd.Username == "" || cmd.Email == "" || cmd.Password == "" {
		return 0, errors.New("missing required fields")
	}

	if cmd.Role != "medium" && cmd.Role != "regular" {
		return 0, errors.New("role must be 'medium' or 'regular'")
	}

	// ДОБАВИТЬ: Хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, errors.New("failed to hash password")
	}

	u := &entities.User{
		Username:   cmd.Username,
		Surname:    cmd.Surname,
		Patronymic: cmd.Patronymic,
		Email:      strings.ToLower(cmd.Email),
		Password:   string(hashedPassword), // Использовать хешированный пароль
		Role:       cmd.Role,
	}

	if err := h.Users.Create(ctx, u); err != nil {
		return 0, err
	}

	return u.ID, nil
}
