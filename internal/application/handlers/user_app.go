package handlers

import (
	"context"
	"taskqueue/internal/application/commands"
	"taskqueue/internal/domain/repositories"
)

type UserApp struct {
	CreateUser *commands.CreateUserHandler
}

func NewUserApp(userRepo repositories.UserRepository) *UserApp {
	return &UserApp{
		CreateUser: commands.NewCreateUserHandler(userRepo),
	}
}

func (app *UserApp) CreateUserBySuper(ctx context.Context, username, surname, patronymic, email, password, role string) (int, error) {
	cmd := commands.CreateUser{
		Username:   username,
		Surname:    surname,
		Patronymic: patronymic,
		Email:      email,
		Password:   password,
		Role:       role,
	}

	return app.CreateUser.Handle(ctx, cmd)
}
