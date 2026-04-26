package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response represents a standard API response
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Success returns a standard success response
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

// Error returns a standard error response
func Error(c *gin.Context, httpCode int, errCode int, message string) {
	c.JSON(httpCode, Response{
		Code:    errCode,
		Message: message,
	})
}
