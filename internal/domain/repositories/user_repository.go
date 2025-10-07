package repositories

import (
	"context"
	"taskqueue/internal/domain/entities"
)

type UserRepository interface {
	Create(ctx context.Context, user *entities.User) error
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
	FindByUserID(ctx context.Context, id int) (*entities.User, error)
	GetAll(ctx context.Context) ([]*entities.User, error)
}
