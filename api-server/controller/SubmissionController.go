package controller

import (
	"context"
	"log"
	"ustoj-master/common"
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
	/*	var user, u model.User
		code := vo.OK*/

	/*defer func() {
		resp := vo.RegisterResponse{
			Code: code,
		}
		c.JSON(http.StatusOK, resp)
		utils.LogReqRespBody(req, resp, "CreateMember")
	}()*/

	if err := c.ShouldBindJSON(&req); err != nil {

		log.Println("CreateMember: ShouldBindJSON error")
		return
	}

	//User existed

	log.Println("CreateMember:UserExisted")
	return
}
