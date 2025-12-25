package controller

import (
	"github.com/gin-gonic/gin"
)

func (uc *UserController) PublicRoutes(router *gin.Engine) {
	routerGroup := router.Group("/user")
	routerGroup.GET("list", uc.list)
}
