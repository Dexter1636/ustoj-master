package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"ustoj-master/common"
	"ustoj-master/model"
	"ustoj-master/service"
	"ustoj-master/vo"

	"github.com/dgrijalva/jwt-go"
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

//var logger = common.LogInstance()

func NewProblemController() IProblemController { // Similar to the interface of service
	//return ProblemController{DB: common.GetDB(), Ctx: common.GetCtx()}
	return ProblemController{DB: common.GetDB(), Ctx: common.GetCtx()}
}

func (ctl ProblemController) ProblemList(c *gin.Context) {
	code := vo.OK
	var problemlist []model.Problem
	DBService := service.NewDBConnect()
	JWTService := service.NewJWTService()
	var logger = common.LogInstance()
	var Username = ""
	defer func() {
		resp := vo.ProblemListResponse{
			Code:        code,
			Problemlist: problemlist,
			Username:    Username,
		}
		c.JSON(http.StatusOK, resp)
		//utils.LogReqRespBody(req, resp, "ReturnProblemPage")
	}()
	var req vo.ProblemListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		code = vo.UnknownError
		//log.Print("ProblemList: BindQuery error")
		logger.Infoln("ProblemList: BindQuery error")
		return
	}
	logger.Infoln("ProblemList: page size: %+v\n", req.Page_Size)
	problemlist = DBService.GetProblemList(problemlist)
	autoHeader := c.GetHeader("Authorization")
	token, errToken := JWTService.ValidateToken(autoHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	//name, err := strconv.ParseUint(fmt.Sprintf("%v", claims["Username"]), 10, 64)
	name := fmt.Sprintf("%v", claims["Username"])
	log.Print(name)

	Username = name
	return
}

func (ctl ProblemController) ProblemDetail(c *gin.Context) {

	var req vo.ProblemDetailRequest
	var problem, p model.Problem
	//	var descriptionModel, d model.Description
	var d model.Description
	code := vo.OK
	var problemID = 0
	var description = ""
	var status = ""
	var difficulty = ""
	var acceptance = ""
	var globalAcceptance = ""
	DBService := service.NewDBConnect()
	JWTService := service.NewJWTService()
	var logger = common.LogInstance()
	var Username = ""
	defer func() {
		resp := vo.ProblemDetailResponse{
			Code:              code,
			ProblemID:         problemID,
			Description:       description,
			Status:            status,
			Difficulty:        difficulty,
			Acceptance:        acceptance,
			Global_Acceptance: globalAcceptance,
			Username:          Username,
		}
		c.JSON(http.StatusOK, resp)
		//utils.LogReqRespBody(req, resp, "ReturnProblemDescription")
	}()
	if err := c.BindQuery(&req); err != nil {
		code = vo.UnknownError
		logger.Infoln("ProblemList: BindQuery error")
		return
	}
	problem.ProblemID = req.ProblemID
	p = DBService.ProblemDetail(req.ProblemID)
	problemID = p.ProblemID
	status = p.Status
	difficulty = p.Difficulty
	acceptance = p.Acceptance
	globalAcceptance = p.GlobalAcceptance

	d = DBService.ProblemDescription(req.ProblemID)
	description = string(d.Description)
	autoHeader := c.GetHeader("Authorization")
	token, errToken := JWTService.ValidateToken(autoHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)

	name := fmt.Sprintf("%v", claims["Username"])

	Username = name
	return

}
