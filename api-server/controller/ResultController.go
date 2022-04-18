package controller

import (
	"context"
	"log"
	"ustoj-master/common"
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
	/*var user, u model.User
	code := vo.OK*/

	/*defer func() {
		resp := vo.RegisterResponse{
			Code: code,
		}
		c.JSON(http.StatusOK, resp)
		utils.LogReqRespBody(req, resp, "CreateMember")
	}()*/

	if err := c.ShouldBindJSON(&req); err != nil {
		//code = vo.UnknownError
		log.Println("CreateMember: ShouldBindJSON error")
		return
	}

	return
}
