package dto

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
