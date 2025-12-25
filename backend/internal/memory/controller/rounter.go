package controller

import "github.com/gin-gonic/gin"

func (mc *MemoryController) PublicRoutes(router *gin.Engine) {
	routerGroup := router.Group("/memory")
	routerGroup.GET("list", mc.list)
}
