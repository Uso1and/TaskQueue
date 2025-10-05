package routes

import (
	"taskqueue/internal/application/handlers"
	"taskqueue/internal/infrastructure/persistence/postgresql"
	httpHandlers "taskqueue/internal/interfaces/http/handlers"
	"taskqueue/internal/interfaces/http/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	userRepo := postgresql.NewUserRepository(postgresql.DB)

	userApp := handlers.NewUserApp(userRepo)

	userHTTPHandler := httpHandlers.NewUserHTTPHandler(userApp)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{

		users := v1.Group("/users") //UserCreate
		{

			users.POST("", middleware.RequireSuper(), userHTTPHandler.CreateUser)
		}
	}

	return r
}
