package controller

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"regexp"
	"ustoj-master/common"
	"ustoj-master/model"
	"ustoj-master/utils"
	"ustoj-master/vo"
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

func NewUserController() IUserController {
	return UserController{DB: common.GetDB(), Ctx: common.GetCtx()}
}

func (ctl UserController) Register(c *gin.Context) {
	var req vo.RegisterRequest
	var user, u model.User
	code := vo.OK

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

	//参数校验
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

	if err := ctl.DB.Where("user_name = ?", req.Username).Take(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctl.DB.Create(&user)
			log.Println("CreateMember:Successfully create, username:" + user.Username)
			return
		} else {
			code = vo.UnknownError
			log.Println("CreateMember:Unknown-error while creating")
			return
		}
	}

	//用户已经存在
	code = vo.UserHasExisted
	log.Println("CreateMember:UserExisted")
	return
}

func (ctl UserController) Login(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (ctl UserController) Logout(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
