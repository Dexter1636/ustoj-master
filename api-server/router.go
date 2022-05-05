package main

import (
	"ustoj-master/api-server/controller"
	"ustoj-master/middleware"
	"ustoj-master/service"

	"github.com/gin-gonic/gin"
)

var (
	jwtService service.JWTService = service.NewJWTService()
)

func RegisterRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Recovery())

	g := r.Group("/api/v1")

	// ping test
	g.GET("/ping", controller.Ping)

	// user
	uc := controller.NewUserController()
	authRoutes := r.Group("/api/vi")
	{
		authRoutes.POST("/user/register", uc.Register)
		authRoutes.POST("/user/login", uc.Login)
		authRoutes.POST("/user/logout", uc.Logout)
	}

	return r
}

func ProblemRouter() *gin.Engine {

	r := gin.Default()

	r.Use(middleware.Recovery())

	p := r.Group("/api/v1")

	// ping test
	p.GET("/ping", controller.Ping)

	// problem
	pc := controller.NewProblemController()
	ProblemRoutes := r.Group("api/v1", middleware.AuthorizenJWT(jwtService))
	{
		ProblemRoutes.POST("/user/problem_list", pc.ProblemList)
		ProblemRoutes.POST("/user/problem_detail", pc.ProblemDetail)
	}
	return r
}

func SubmissionRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Recovery())

	s := r.Group("/api/v1")

	// ping test
	s.GET("/ping", controller.Ping)

	// problem
	sc := controller.NewSubmissionController()
	SubmissionRoutes := r.Group("api/v1", middleware.AuthorizenJWT(jwtService))
	{
		SubmissionRoutes.POST("/user/submit", sc.Submit)
	}

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
	ResultRoutes := r.Group("api/v1", middleware.AuthorizenJWT(jwtService))
	{
		ResultRoutes.POST("/user/result_list", rc.ResultList)
	}
	return r
}
