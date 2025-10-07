package handlers

import (
	"net/http"
	"taskqueue/internal/application/dto"
	"taskqueue/internal/application/handlers"
	"taskqueue/internal/domain/entities"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHTTPHandler struct {
	userApp *handlers.UserApp
}

func NewUserHTTPHandler(userApp *handlers.UserApp) *UserHTTPHandler {
	return &UserHTTPHandler{userApp: userApp}
}

func (h *UserHTTPHandler) CreateUser(c *gin.Context) {

	var req dto.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid request format",
			"details": err.Error(),
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to hash password",
		})
		return
	}

	userID, err := h.userApp.CreateUserBySuper(
		c.Request.Context(),
		req.Username,
		req.Surname,
		req.Patronymic,
		req.Email,
		string(hashedPassword),
		req.Role,
	)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error":   "failed to create user",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.CreateUserResponse{
		ID:      userID,
		Message: "User created successfully",
	})
}

func (h *UserHTTPHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userApp.GetAllUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "failed to get users",
			"details": err.Error(),
		})
		return
	}

	// Конвертируем entities в DTO
	userResponses := make([]dto.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = entityToUserResponse(user)
	}

	c.JSON(http.StatusOK, dto.GetAllUsersResponse{
		Users: userResponses,
		Total: len(userResponses),
	})
}

// Вспомогательная функция для конвертации entity в DTO
func entityToUserResponse(user *entities.User) dto.UserResponse {
	return dto.UserResponse{
		ID:         user.ID,
		Username:   user.Username,
		Surname:    user.Surname,
		Patronymic: user.Patronymic,
		Email:      user.Email,
		Role:       user.Role,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}
}
