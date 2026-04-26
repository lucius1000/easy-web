package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/user/go-gin-gorm-starter/pkg/logger"
	"github.com/user/go-gin-gorm-starter/pkg/response"
	"go.uber.org/zap"
)

func GlobalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			logger.Log.Error("global error", zap.Error(err))
			response.Error(c, http.StatusInternalServerError, 500, "Internal Server Error")
		}
	}
}

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Log.Error("panic recovered", zap.Any("error", err))
				response.Error(c, http.StatusInternalServerError, 500, "Internal Server Error")
				c.Abort()
			}
		}()
		c.Next()
	}
}
