package controller

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
	"ustoj-master/common"
	"ustoj-master/model"
	"ustoj-master/service"
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
	if err := c.BindQuery(&req); err != nil {
		log.Println("ProblemList: BindQuery error")
		return
	} else {
		log.Printf("language: %v\n Code: %v\n", req.Language, req.Code)
	}
	authHeader := c.GetHeader("Authorization")
	token, err := JWTService.ValidateToken(authHeader)
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		log.Println("Claim[issuer]:", claims["issuer"])
	} else {
		log.Print(err)
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
	}
	DBService.Submission(&submission)
	return

}
