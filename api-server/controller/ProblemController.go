package controller

import (
	"context"
	"ustoj-master/common"
	"ustoj-master/model"
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
	//return ProblemController{DB: common.GetDB(), Ctx: common.GetCtx()}
	return ProblemController{DB: common.GetDB(), Ctx: common.GetCtx()}
}

func (ctl ProblemController) ProblemList(c *gin.Context) {
	var req vo.ProblemListRequest
	var problem model.Problem

	/*defer func() {
		resp := vo.ProblemListResponse{
			ProblemID:         problemID,
			Status:            status,
			Difficulty:        difficulty,
			Acceptance:        acceptance,
			Global_Acceptance: globalAcceptance,
		}
		c.JSON(http.StatusOK, resp)
		utils.LogReqRespBody(req, resp, "CreateMember")
	}()*/

	var Page_size = req.Page_Size
	println(Page_size)
	ctl.DB.Create(&problem)

	return
}

func (ctl ProblemController) ProblemDetail(c *gin.Context) {
	//TODO implement me
	panic("implement me")
	var req vo.ProblemDetailRequest
	var problem model.Problem
	var description model.Description

	description = model.Description{ProblemID: req.ProblemID}
	ctl.DB.Create(&description)
	ctl.DB.Create(&problem)
	return
}
