package controller

import (
	"context"
	"ustoj-master/common"
	"ustoj-master/model"
	"ustoj-master/service"
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
	var logger = common.LogInstance()
	DBService := service.NewDBConnect()
	JWTService := service.NewJWTService()
	submission = model.Submission{ProblemID: req.ProblemID, Code: req.Code, Language: req.Language}
	DBService.Submission(&submission)
	autoHeader := c.GetHeader("Authorization")
	token, errToken := JWTService.ValidateToken(autoHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	logger.Infoln(token)
	return

}
