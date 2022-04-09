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

	return r
}
