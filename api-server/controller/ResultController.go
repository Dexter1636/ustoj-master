package controller

import (
	"context"
	"errors"
	"log"
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
	var req vo.ResultRequest
	var user, u model.User
	code := vo.OK

	/*defer func() {
		resp := vo.RegisterResponse{
			Code: code,
		}
		c.JSON(http.StatusOK, resp)
		utils.LogReqRespBody(req, resp, "CreateMember")
	}()*/

	if err := c.ShouldBindJSON(&req); err != nil {
		code = vo.UnknownError
		log.Println("CreateMember: ShouldBindJSON error")
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

	//User existed
	code = vo.UserHasExisted
	log.Println("CreateMember:UserExisted")
	return
}

func (ctl UserController) Logout(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
