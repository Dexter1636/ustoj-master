package controller

import (
	"context"
	"fmt"
	"net/http"
	"ustoj-master/common"
	"ustoj-master/model"
	"ustoj-master/service"
	"ustoj-master/utils"
	"ustoj-master/vo"

	"github.com/dgrijalva/jwt-go"

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
	//var problemID = 0
	//var username = ""
	//var status = ""
	//var language = ""
	//var runtime = 0
	records := make([]*model.Submission, 0)
	code := vo.OK
	DBService := service.NewDBConnect()
	JWTService := service.NewJWTService()
	authHeader := c.GetHeader("Authorization")
	token, err := JWTService.ValidateToken(authHeader)
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		req.Username = fmt.Sprintf("%v", claims["Username"])
	} else {
		logger.Print(err)
		code = vo.UnknownError
		resp := vo.SubmissionResponse{
			Code: code,
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
		return
	}
	defer func() {
		resp := vo.ResultResponse{
			Code:    code,
			Records: records,
		}
		c.JSON(http.StatusOK, resp)
		utils.LogReqRespBody(req, resp, "ReturnProblemPage")
	}()
	if err := c.BindQuery(&req); err != nil {
		code = vo.UnknownError
		logger.Println("ProblemList: BindQuery error")
		return
	}
	submission = model.Submission{ProblemID: req.ProblemID, Username: req.Username}
	s := DBService.ResultList(&submission)
	for i, _ := range s {
		s[i].Code = ""
		records = append(records, &s[i])
	}
	//problemID = s[0].ProblemID
	//username = s[0].Username
	//status = s[0].Status
	//language = s[0].Language
	//runtime = s[0].RunTime
	//logger.Printf("problem_id:%v,language:%v", problemID, language)

	return
}
