package main

import (
	"backend/internal"
	"backend/internal/user/controller"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {
	g := gin.Default()
	app := fx.New(internal.Module, fx.Invoke(func(UC *controller.UserController) {
		UC.PublicRoutes(g)
	}))
	app.Run()
}

func NewGinEngine() *gin.Engine {
	r := gin.Default()
	return r
}
