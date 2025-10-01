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
	query := `INSERT INTO users(username, age, email, created_at) VALUES ($1, $2, $3, $4) RETURNING id`

	return ur.db.QueryRowContext(ctx, query, user.Username, user.Age, user.Email, user.CreatedAt).Scan(&user.ID)
}
