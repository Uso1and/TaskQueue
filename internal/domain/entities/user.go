package entities

import "time"

type SyperUser struct {
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	Surname    string    `json:"surname"`
	Patronymic string    `json:"patronymic"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	Role       string    `json:"role"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
