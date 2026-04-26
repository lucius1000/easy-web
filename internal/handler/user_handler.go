package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/user/go-gin-gorm-starter/internal/service"
	"github.com/user/go-gin-gorm-starter/pkg/response"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, 400, err.Error())
		return
	}

	user, err := h.userService.CreateUser(c.Request.Context(), req.Name, req.Email, req.Password)
	if err != nil {
		if err == service.ErrEmailAlreadyExists {
			response.Error(c, http.StatusConflict, 409, err.Error())
			return
		}
		response.Error(c, http.StatusInternalServerError, 500, "Failed to create user")
		return
	}

	response.Success(c, user)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, 400, "Invalid user ID")
		return
	}

	user, err := h.userService.GetUserByID(c.Request.Context(), uint(id))
	if err != nil {
		if err == service.ErrUserNotFound {
			response.Error(c, http.StatusNotFound, 404, err.Error())
			return
		}
		response.Error(c, http.StatusInternalServerError, 500, "Failed to get user")
		return
	}

	response.Success(c, user)
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")

	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	users, total, err := h.userService.ListUsers(c.Request.Context(), page, pageSize)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, 500, "Failed to list users")
		return
	}

	response.Success(c, gin.H{
		"items": users,
		"total": total,
	})
}

type UpdateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email" binding:"omitempty,email"`
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, 400, "Invalid user ID")
		return
	}

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, 400, err.Error())
		return
	}

	user, err := h.userService.UpdateUser(c.Request.Context(), uint(id), req.Name, req.Email)
	if err != nil {
		if err == service.ErrUserNotFound {
			response.Error(c, http.StatusNotFound, 404, err.Error())
			return
		}
		if err == service.ErrEmailAlreadyExists {
			response.Error(c, http.StatusConflict, 409, err.Error())
			return
		}
		response.Error(c, http.StatusInternalServerError, 500, "Failed to update user")
		return
	}

	response.Success(c, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, 400, "Invalid user ID")
		return
	}

	err = h.userService.DeleteUser(c.Request.Context(), uint(id))
	if err != nil {
		if err == service.ErrUserNotFound {
			response.Error(c, http.StatusNotFound, 404, err.Error())
			return
		}
		response.Error(c, http.StatusInternalServerError, 500, "Failed to delete user")
		return
	}

	response.Success(c, nil)
}
