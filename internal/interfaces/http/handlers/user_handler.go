package handlers

import (
	"log"
	"net/http"
	"taskqueue/internal/domain/entities"
	"taskqueue/internal/domain/repositories"

	"time"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userRepo repositories.UserRepository
}

func NewUserHandelr(userRepo repositories.UserRepository) *UserHandler {
	return &UserHandler{userRepo: userRepo}
}

func (uh *UserHandler) CreateUserHandler(ctx *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Age      int    `json:"age"`
		Email    string `json:"email"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if req.Username == "" || req.Email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "required"})
		return
	}

	user := entities.User{
		Username:  req.Username,
		Age:       req.Age,
		Email:     req.Email,
		CreatedAt: time.Now(),
	}

	if err := uh.userRepo.CreateUserRepo(ctx.Request.Context(), &user); err != nil {
		log.Printf("error create user: %v", err)

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Create user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User succesfully create",
		"user":    user,
	})
}
