package routes

import (
	"taskqueue/internal/infrastructure/persistence/postgresql"
	"taskqueue/internal/interfaces/http/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	userRepo := postgresql.NewSuperUserRepo(postgresql.DB)

	superUserHandler := handlers.NewSuperUserHandelr(userRepo)

	r := gin.Default()

	r.POST("/sucrt", superUserHandler.CreateSuperUserHandler)
	return r
}
