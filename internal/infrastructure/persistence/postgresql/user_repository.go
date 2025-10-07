package postgresql

import (
	"context"
	"database/sql"
	"taskqueue/internal/domain/entities"
	"taskqueue/internal/domain/repositories"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repositories.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *entities.User) error {

	const q = `
	INSERT INTO users (username, surname, patronymic, password, email, role)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id, created_at, updated_at
	`

	return r.db.QueryRowContext(ctx, q,
		user.Username,
		user.Surname,
		user.Patronymic,
		user.Password,
		user.Email,
		user.Role,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {

	const q = `
	SELECT id, username, surname, patronymic, password, email, role, created_at, updated_at FROM users WHERE email = $1
	`

	user := &entities.User{}

	err := r.db.QueryRowContext(ctx, q, email).Scan(
		&user.ID, &user.Username, &user.Surname, &user.Patronymic, &user.Password, &user.Email,
		&user.Role, &user.CreatedAt, &user.UpdatedAt,
	)

	return user, err
}

func (r *UserRepository) FindByUserID(ctx context.Context, id int) (*entities.User, error) {

	const q = `
	SELECT id, username, surname, patronymic, password, email, role, created_at, updated_at FROM users WHERE id = $1
	`

	user := &entities.User{}

	err := r.db.QueryRowContext(ctx, q, id).Scan(
		&user.ID, &user.Username, &user.Surname, &user.Patronymic, &user.Password, &user.Email,
		&user.Role, &user.Role, &user.CreatedAt, &user.UpdatedAt,
	)

	return user, err
}

func (r *UserRepository) GetAll(ctx context.Context) ([]*entities.User, error) {

	const q = `
	SELECT id, username, surname, patronymic, password, email, role, created_at, updated_at
	FROM users ORDER BY created_at DESC`

	rows, err := r.db.QueryContext(ctx, q)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*entities.User

	for rows.Next() {

		user := &entities.User{}

		err := rows.Scan(
			&user.ID, &user.Username, &user.Surname, &user.Patronymic, &user.Password,
			&user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, err
}
