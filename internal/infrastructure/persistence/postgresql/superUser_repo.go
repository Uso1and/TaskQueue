package postgresql

import (
	"context"
	"database/sql"
	"taskqueue/internal/domain/entities"
	"taskqueue/internal/domain/repositories"
)

type PostSuperUserRepo struct {
	db *sql.DB
}

func NewSuperUserRepo(db *sql.DB) repositories.SuperUserRepository {
	return &PostSuperUserRepo{db: db}
}

func (sur *PostSuperUserRepo) CreateSuperUserRepo(ctx context.Context, user *entities.SyperUser) error {
	const q = `
        INSERT INTO superusers (username, surname, patronymic, password, email, role)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id
    `
	return sur.db.QueryRowContext(
		ctx, q,
		user.Username, user.Surname, user.Patronymic, user.Password, user.Email, user.Role,
	).Scan(&user.ID)
}
