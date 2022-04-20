package controller

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"
	"ustoj-master/common"
	"ustoj-master/model"
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

	submission = model.Submission{ProblemID: req.ProblemID, Username: req.Username}
	if err := ctl.DB.Where("problem_id =?", req.ProblemID).Where("username = ?", req.Username).Take(&s).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctl.DB.First(&submission)
			log.Println("Successfully find the problem result" + strconv.Itoa(submission.ProblemID))
			problemID = submission.ProblemID
			username = submission.Username
			status = submission.Status
			language = submission.Language
			runtime = submission.RunTime
			return
		} else {
			code = vo.UnknownError
			log.Println("Result :Unknown-error while finding Problem information")
			return
		}
	}
	return
}
