package controller

import (
	"backend/internal/user/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (UC *UserController) list(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "user list",
	})
}
