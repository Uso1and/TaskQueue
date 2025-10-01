package repositories

import (
	"context"
	"taskqueue/internal/domain/entities"
)

type UserRepository interface {
	CreateUserRepo(ctx context.Context, user *entities.User) error
}
