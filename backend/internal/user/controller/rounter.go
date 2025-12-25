package controller

import (
	"github.com/gin-gonic/gin"
)

func (UC *UserController) PublicRoutes(router *gin.Engine) {
	routerGroup := router.Group("/user")
	routerGroup.GET("list", UC.list)
}
