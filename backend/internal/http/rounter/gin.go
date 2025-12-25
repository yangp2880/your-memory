package rounter

import (
	mc "backend/internal/memory/controller"
	uc "backend/internal/user/controller"

	"github.com/gin-gonic/gin"
)

func NewGinEngine(uc *uc.UserController, mc *mc.MemoryController) *gin.Engine {
	r := gin.Default()
	uc.PublicRoutes(r)
	mc.PublicRoutes(r)
	return r
}
