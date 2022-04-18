package main

import (
	"ustoj-master/api-server/controller"
	"ustoj-master/middleware"

	"github.com/gin-gonic/gin"
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

func ProblemRouter() *gin.Engine {

	r := gin.Default()

	r.Use(middleware.Recovery())

	g := r.Group("/api/v1")

	// ping test
	g.GET("/ping", controller.Ping)

	// problem
	pc := controller.NewProblemController()
	g.POST("/user/problem_list", pc.ProblemList)
	g.POST("/user/problem_detail", pc.ProblemDetail)

	return r
}

func SubmissionRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Recovery())

	g := r.Group("/api/v1")

	// ping test
	g.GET("/ping", controller.Ping)

	// problem
	sc := controller.NewSubmissionController()
	g.POST("/user/submit", sc.Submit)

	return r
}
func ResultRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Recovery())

	g := r.Group("/api/v1")

	// ping test
	g.GET("/ping", controller.Ping)

	// problem
	rc := controller.NewResultController()
	g.POST("/user/result_list", rc.ResultList)

	return r
}
