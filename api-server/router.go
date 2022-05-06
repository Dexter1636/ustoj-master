package main

import (
	"net/http"
	"ustoj-master/api-server/controller"
	"ustoj-master/middleware"
	"ustoj-master/service"

	"github.com/gin-gonic/gin"
)

var (
	jwtService service.JWTService = service.NewJWTService()
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func RegisterRouter() *gin.Engine {
	r := gin.Default()
	r.Use(Cors())

	r.Use(middleware.Recovery())

	g := r.Group("/api/v1")

	// ping test
	g.GET("/ping", controller.Ping)

	// user
	uc := controller.NewUserController()
	authRoutes := r.Group("/api/v1")
	{
		authRoutes.POST("/user/register", uc.Register)
		authRoutes.POST("/user/login", uc.Login)
		authRoutes.POST("/user/logout", uc.Logout)
	}

	pc := controller.NewProblemController()
	ProblemRoutes := r.Group("/api/v1", middleware.AuthorizenJWT(jwtService))
	{
		ProblemRoutes.GET("/user/problem_list", pc.ProblemList)
		ProblemRoutes.GET("/user/problem_detail", pc.ProblemDetail)
	}
	sc := controller.NewSubmissionController()
	SubmissionRoutes := r.Group("/api/v1", middleware.AuthorizenJWT(jwtService))
	{
		SubmissionRoutes.POST("/user/submit", sc.Submit)
	}
	rc := controller.NewResultController()
	ResultRoutes := r.Group("/api/v1", middleware.AuthorizenJWT(jwtService))
	{
		ResultRoutes.GET("/user/result_list", rc.ResultList)
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
		ProblemRoutes.GET("/user/problem_list", pc.ProblemList)
		ProblemRoutes.GET("/user/problem_detail", pc.ProblemDetail)
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
		SubmissionRoutes.GET("/user/submit", sc.Submit)
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
		ResultRoutes.GET("/user/result_list", rc.ResultList)
	}
	return r
}
