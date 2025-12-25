package controller

import (
	"backend/internal/user/service"

	"github.com/gin-gonic/gin"
)

type MemoryController struct {
	service service.UserService
}

func NewMemoryController(service service.UserService) *MemoryController {
	return &MemoryController{
		service: service,
	}
}

func (mc *MemoryController) list(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "memory list",
	})
}
