package controller

import (
	"context"
	//"fmt"

	"net/http"

	//	"strconv"
	"ustoj-master/common"
	"ustoj-master/model"
	"ustoj-master/service"
	"ustoj-master/utils"
	"ustoj-master/vo"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var logger = common.LogInstance()

type IUserController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

type UserController struct {
	DB  *gorm.DB
	Ctx context.Context
}

func NewUserController() IUserController { // Similar to the interface of service
	return UserController{DB: common.GetDB(), Ctx: common.GetCtx()}
}

func (ctl UserController) Register(c *gin.Context) {
	var req vo.RegisterRequest
	//var user, u model.User
	var user model.User
	code := vo.OK
	DBService := service.NewDBConnect()
	defer func() {
		resp := vo.RegisterResponse{
			Code: code,
		}
		c.JSON(http.StatusOK, resp)
		utils.LogReqRespBody(req, resp, "CreateMember")
	}()

	if err := c.ShouldBind(&req); err != nil {
		code = vo.UnknownError
		resp := vo.SubmissionResponse{
			Code: code,
		}
		logger.Errorln(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}

	user = model.User{Username: req.Username, Password: req.Password, RoleId: 1}
	logger.Infoln("username:" + user.Username + "password:" + user.Password)
	code, err := DBService.CreateUser(&user)
	if err != nil {
		code = vo.UnknownError
		resp := vo.SubmissionResponse{
			Code: code,
		}
		logger.Errorln(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	return
}

func (ctl UserController) Login(c *gin.Context) {
	var req vo.LoginRequest
	var user model.User
	var loginobject vo.LoginResponse
	DBService := service.NewDBConnect()
	JWTService := service.NewJWTService()
	var Token = ""
	code := vo.OK
	defer func() {
		resp := vo.LoginResponse{
			Code:  code,
			Data:  loginobject.Data,
			Token: Token,
		}
		c.JSON(http.StatusOK, resp)
		//utils.LogReqRespBody(req, resp, "XXXXXXXXXXX")
	}()
	if err := c.ShouldBind(&req); err != nil {
		code = vo.UnknownError
		logger.Println("Login: ShouldBindJSON error")
		return
	}
	user = model.User{Username: req.Username, Password: req.Password, RoleId: 1}
	loginobject.Data.UserID = DBService.Login(&user)
	if loginobject.Data.UserID == "UnknownError" {
		code = vo.UnknownError
		Token = ""
	} else {
		Token = JWTService.GenerateToken(loginobject.Data.UserID)
	}
}

func (ctl UserController) Logout(c *gin.Context) {

}
