package dto

import "time"

type CreateUserRequest struct {
	Username   string `json:"username" binding:"required"`
	Surname    string `json:"surname" binding:"required"`
	Patronymic string `json:"patronymic" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=6"`
	Role       string `json:"role" binding:"required,oneof=medium regular"`
}

type CreateUserResponse struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

// Добавляем недостающие структуры для получения пользователей
type UserResponse struct {
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	Surname    string    `json:"surname"`
	Patronymic string    `json:"patronymic"`
	Email      string    `json:"email"`
	Role       string    `json:"role"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type GetAllUsersResponse struct {
	Users []UserResponse `json:"users"`
	Total int            `json:"total"`
}
