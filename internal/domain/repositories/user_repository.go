package repositories

import (
	"context"
	"taskqueue/internal/domain/entities"
)

type SuperUserRepository interface {
	CreateSuperUserRepo(ctx context.Context, user *entities.SyperUser) error
}
