package handlers

import (
	"net/http"
	"taskqueue/internal/domain/entities"
	"taskqueue/internal/domain/repositories"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type SyperUserHandler struct {
	userRepo repositories.SuperUserRepository
}

func NewSuperUserHandelr(userRepo repositories.SuperUserRepository) *SyperUserHandler {
	return &SyperUserHandler{userRepo: userRepo}
}

func (suh *SyperUserHandler) CreateSuperUserHandler(ctx *gin.Context) {
	var required struct {
		Username   string `json:"username"`
		Surname    string `json:"surname"`
		Patronymic string `json:"patronymic"`
		Password   string `json:"password"`
		Email      string `json:"email"`
		Role       string `json:"role"`
	}
	if err := ctx.ShouldBindJSON(&required); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if required.Username == "" || required.Surname == "" || required.Patronymic == "" ||
		required.Password == "" || required.Email == "" || required.Role == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing fields"})
		return
	}

	hashpassword, err := bcrypt.GenerateFromPassword([]byte(required.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	su := entities.SyperUser{
		Username:   required.Username,
		Surname:    required.Surname,
		Patronymic: required.Patronymic,
		Password:   string(hashpassword),
		Email:      required.Email,
		Role:       required.Role,
		// created_at/updated_at даст БД по DEFAULT
	}

	if err := suh.userRepo.CreateSuperUserRepo(ctx.Request.Context(), &su); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "superUser created",
		"id":      su.ID,
	})
}
