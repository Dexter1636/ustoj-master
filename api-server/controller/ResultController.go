package controller

import (
	"context"
	"log"
	"net/http"
	"ustoj-master/common"
	"ustoj-master/model"
	"ustoj-master/service"
	"ustoj-master/utils"
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
	var submission, s model.Submission
	var problemID = 0
	var username = ""
	var status = ""
	var language = ""
	var runtime = 0
	code := vo.OK
	DBService := service.NewDBConnect()
	JWTService := service.NewJWTService()
	defer func() {
		resp := vo.ResultResponse{
			Code:      code,
			ProblemID: problemID,
			Username:  username,
			Status:    status,
			Language:  language,
			RunTime:   runtime,
		}
		c.JSON(http.StatusOK, resp)
		utils.LogReqRespBody(req, resp, "ReturnProblemPage")
	}()
	if err := c.BindQuery(&req); err != nil {
		code = vo.UnknownError
		log.Println("ProblemList: BindQuery error")
		return
	}
	submission = model.Submission{ProblemID: req.ProblemID, Username: req.Username}
	s = DBService.ResultList(&submission)
	problemID = s.ProblemID
	username = s.Username
	status = s.Status
	language = s.Language
	runtime = s.RunTime

	if problemID == 0 {
		code = vo.UnknownError
	}
	autoHeader := c.GetHeader("Authorization")
	token, errToken := JWTService.ValidateToken(autoHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	print(token)
	return
}
