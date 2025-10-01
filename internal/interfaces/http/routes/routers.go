package routes

import (
	"taskqueue/internal/infrastructure/persistence/postgresql"
	"taskqueue/internal/interfaces/http/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	userRepo := postgresql.NewUserRepo(postgresql.DB)

	userHandler := handlers.NewUserHandelr(userRepo)

	r := gin.Default()

	r.POST("/cruser", userHandler.CreateUserHandler)

	return r
}
