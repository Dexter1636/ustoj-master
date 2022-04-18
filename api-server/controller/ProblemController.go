package controller

import (
	"context"
	"log"
	"ustoj-master/common"
	"ustoj-master/vo"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IProblemController interface {
	ProblemList(c *gin.Context)
	ProblemDetail(c *gin.Context)
}

type ProblemController struct {
	DB  *gorm.DB
	Ctx context.Context
}

func NewProblemController() IProblemController { // Similar to the interface of service
	return ProblemController{DB: common.GetDB(), Ctx: common.GetCtx()}
}

func (ctl ProblemController) ProblemList(c *gin.Context) {
	var req vo.ProblemListRequest
	//var problem, p model.Problem
	//code := vo.OK

	//problem = model.Problem{ProblemID: req.ProblemID, Status: req.Status, RoleId: 1}

	//User existed
	if err := c.ShouldBindJSON(&req); err != nil {
		//code = vo.UnknownError
		log.Println("Login: ShouldBindJSON error")
		return
	}
}
func (ctl ProblemController) ProblemDetail(c *gin.Context) {
	var req vo.ProblemListRequest
	//var problem, p model.Problem
	//code := vo.OK

	//problem = model.Problem{ProblemID: req.ProblemID, Status: req.Status, RoleId: 1}

	//User existed
	if err := c.ShouldBindJSON(&req); err != nil {
		//code = vo.UnknownError
		log.Println("Login: ShouldBindJSON error")
		return
	}
}
