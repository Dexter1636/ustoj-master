package controller

import (
	"context"
	/*"errors"
	"log"*/
	"ustoj-master/common"
	"ustoj-master/model"
	"ustoj-master/vo"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ISubmissionController interface {
	Submit(c *gin.Context)
}

type SubmissionController struct {
	DB  *gorm.DB
	Ctx context.Context
}

func NewSubmissionController() ISubmissionController { // Similar to the interface of service
	return SubmissionController{DB: common.GetDB(), Ctx: common.GetCtx()}
}

func (ctl SubmissionController) Submit(c *gin.Context) {
	var req vo.SubmissionRequest
	var submission model.Submission

	submission = model.Submission{ProblemID: req.ProblemID, Code: req.Code, Language: req.Language}
	ctl.DB.Create(&submission)
	return
}
