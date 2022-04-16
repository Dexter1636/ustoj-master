package controller

import (
	"context"
	"ustoj-master/common"
	"ustoj-master/model"
	"ustoj-master/vo"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IResultController interface {
	ResultList(c *gin.Context)
}

type ResultController struct {
	DB  *gorm.DB
	Ctx context.Context
}

func NewResultController() IResultController { // Similar to the interface of service
	return ResultController{DB: common.GetDB(), Ctx: common.GetCtx()}
}

func (ctl ResultController) ResultList(c *gin.Context) {
	var req vo.ResultRequest //This varaible controls the request  json object
	var submission model.Submission

	submission = model.Submission{ProblemID: req.ProblemID}
	ctl.DB.Create(&submission)
	return
}
