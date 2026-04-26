package router

import (
	"github.com/gin-gonic/gin"
	"github.com/user/go-gin-gorm-starter/internal/handler"
	"github.com/user/go-gin-gorm-starter/internal/middleware"
	"github.com/user/go-gin-gorm-starter/pkg/response"
)

func SetupRouter(userHandler *handler.UserHandler) *gin.Engine {
	r := gin.New()

	// Middlewares
	r.Use(middleware.ZapLogger())
	r.Use(middleware.Recovery())
	r.Use(middleware.GlobalErrorHandler())

	// Health Check
	r.GET("/health", func(c *gin.Context) {
		response.Success(c, "OK")
	})

	// API Routes
	api := r.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.POST("", userHandler.CreateUser)
			users.GET("", userHandler.ListUsers)
			users.GET("/:id", userHandler.GetUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}
	}

	return r
}
