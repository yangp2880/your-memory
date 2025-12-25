package http

import "github.com/gin-gonic/gin"

func NewGinEngine() *gin.Engine {
	r := gin.Default()
	return r
}
