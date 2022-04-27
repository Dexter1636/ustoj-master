package controller

import (
	"context"
	"log"
	"net/http"
	"regexp"
	"ustoj-master/common"
	"ustoj-master/model"
	"ustoj-master/service"
	"ustoj-master/utils"
	"ustoj-master/vo"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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
	var DBService service.DBService

	defer func() {
		resp := vo.RegisterResponse{
			Code: code,
		}
		c.JSON(http.StatusOK, resp)
		utils.LogReqRespBody(req, resp, "CreateMember")
	}()

	if err := c.ShouldBindJSON(&req); err != nil {
		code = vo.UnknownError
		log.Println("CreateMember: ShouldBindJSON error")
		return
	}

	//Parameter validation
	tmpStr := req.Password
	r1, _ := regexp.MatchString("^(\\w*[0-9]+\\w*[a-z]+\\w*[A-Z]+)|(\\w*[0-9]+\\w*[A-Z]+\\w*[a-z]+)$", tmpStr)
	r2, _ := regexp.MatchString("^(\\w*[a-z]+\\w*[0-9]+\\w*[A-Z]+)|(\\w*[a-z]+\\w*[A-Z]+\\w*[0-9]+)$", tmpStr)
	r3, _ := regexp.MatchString("^(\\w*[A-Z]+\\w*[a-z]+\\w*[0-9]+)|(\\w*[A-Z]+\\w*[0-9]+\\w*[a-z]+)$", tmpStr)
	ru, _ := regexp.MatchString("^([a-z]|[A-Z])*$", req.Username)
	rp := r1 || r2 || r3

	if (len(req.Password) > 20 || len(req.Password) < 8 || !rp) ||
		(len(req.Username) < 8 || len(req.Username) > 20 || !ru) {
		code = vo.ParamInvalid
		log.Println("CreateMember:ParamInvalid")
		return
	}

	user = model.User{Username: req.Username, Password: req.Password, RoleId: 1}
	code = DBService.CreateUser(&user)
	/*if err := ctl.DB.Where("user_name = ?", req.Username).Take(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctl.DB.Create(&user)
			log.Println("CreateMember:Successfully create, username:" + user.Username)
			return
		} else {
			code = vo.UnknownError
			log.Println("CreateMember:Unknown-error while creating")
			return
		}
	}*/

	//User existed
	code = vo.UserHasExisted
	log.Println("CreateMember:UserExisted")
	return
}

func (ctl UserController) Login(c *gin.Context) {
	//TODO implement me
	panic("implement me")
	var req vo.LoginRequest
	var user model.User
	var loginobject vo.LoginResponse
	var DBService service.DBService
	code := vo.OK
	defer func() {
		resp := vo.LoginResponse{
			Code: code,
			Data: loginobject.Data,
		}
		c.JSON(http.StatusOK, resp)
		utils.LogReqRespBody(req, resp, "XXXXXXXXXXX")
	}()
	if err := c.ShouldBindJSON(&req); err != nil {
		code = vo.UnknownError
		log.Println("Login: ShouldBindJSON error")
		return
	}
	user = model.User{Username: req.Username, Password: req.Password, RoleId: 1}
	loginobject.Data.UserID = DBService.Login(&user)
	if loginobject.Data.UserID == "" {
		code = vo.UnknownError
	}
	/*	if err := ctl.DB.Where("user_name = ?", req.Username).Where("password =?", req.Password).Take(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			loginobject.Data.UserID = u.Username
		} else {
			code = vo.UnknownError
			log.Println("Not User record")
			return
		}
	}*/

}

func (ctl UserController) Logout(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
