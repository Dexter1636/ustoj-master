package main

import (
	"github.com/gin-gonic/gin"
	"ustoj-master/api-server/controller"
	"ustoj-master/middleware"
)

func RegisterRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Recovery())

	g := r.Group("/api/v1")

	// ping test
	g.GET("/ping", controller.Ping)

	// user
	uc := controller.NewUserController()
	g.POST("/user/create", uc.Register)
	g.POST("/user/login", uc.Login)
	g.POST("/user/logout", uc.Logout)

	return r
}
