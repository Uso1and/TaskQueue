package postgresql

import (
	"context"
	"database/sql"
	"taskqueue/internal/domain/entities"
	"taskqueue/internal/domain/repositories"
)

type PostUserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) repositories.UserRepository {
	return &PostUserRepo{db: db}
}

func (ur *PostUserRepo) CreateUserRepo(ctx context.Context, user *entities.User) error {
	query := `INSERT INTO users (email, username, password_hash, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	return ur.db.QueryRowContext(ctx, query, user.Email, user.Username, user.Password, user.CreatedAt, user.UpdatedAt).Scan(&user.ID)
}
