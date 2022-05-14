package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"ustoj-master/common"
	"ustoj-master/model"
	"ustoj-master/service"
	"ustoj-master/vo"

	"github.com/dgrijalva/jwt-go"

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
	var submission model.Submission
	DBService := service.NewDBConnect()
	JWTService := service.NewJWTService()
	code := vo.OK
	defer func() {
		resp := vo.SubmissionResponse{
			Code: code,
		}
		c.JSON(http.StatusOK, resp)
	}()
	if err := c.ShouldBind(&req); err != nil {
		logger.Println("Submit: BindQuery error: " + err.Error())
		code = vo.UnknownError
		return
	} else {
		logger.Printf("language: %v\n Code: %v\n", req.Language, req.Code)
	}
	authHeader := c.GetHeader("Authorization")
	token, err := JWTService.ValidateToken(authHeader)
	if err != nil {
		code = vo.UnknownError
		logger.Println("Submit: ValidateToken Error:" + err.Error())
		return
	}
	username := ""
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		username = fmt.Sprintf("%v", claims["Username"])
	} else {
		code = vo.UnknownError
		resp := vo.SubmissionResponse{
			Code: code,
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
		return
	}

	submission = model.Submission{
		SubmissionTime: time.Now(),
		ProblemID:      req.ProblemID,
		Code:           req.Code,
		Language:       req.Language,
		Username:       username,
		Status:         "submitted",
	}
	DBService.Submission(&submission)
	return

}
