package controller

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"
	"ustoj-master/common"
	"ustoj-master/model"
	"ustoj-master/utils"
	"ustoj-master/vo"

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

func NewProblemController() IProblemController { // Similar to the interface of service
	//return ProblemController{DB: common.GetDB(), Ctx: common.GetCtx()}
	return ProblemController{DB: common.GetDB(), Ctx: common.GetCtx()}
}

func (ctl ProblemController) ProblemList(c *gin.Context) {
	var req vo.ProblemListRequest
	var problem []model.Problem
	code := vo.OK
	/*	var problemID = 0
		var status = ""
		var difficulty = ""
		var acceptance = ""
		var globalAcceptance = ""*/
	var problemlist []model.Problem

	defer func() {
		resp := vo.ProblemListResponse{
			Code: code,
			/*ProblemID:         problemID,
			Status:            status,
			Difficulty:        difficulty,
			Acceptance:        acceptance,
			Global_Acceptance: globalAcceptance,*/
			Problemlist: problemlist,
		}
		c.JSON(http.StatusOK, resp)
		utils.LogReqRespBody(req, resp, "ReturnProblemPage")
	}()

	var Page_size = req.Page_Size
	println(Page_size)

	if result := ctl.DB.Find(&problem); result.Error != nil {
		log.Println("Error occured during get all problem information! ")
		return
	} else {

		log.Println("The lenght of all problem :" + strconv.FormatInt(result.RowsAffected, 10))
		/*result should be a slice*/
		/*	for problemdetail := range problem {
			modelList = append(modelList, model.Problem{
				problemID:        problemdetail.ProblemID,
				status:           problemdetail.Status,
				difficulty:       problemdetail.Difficulty,
				acceptance:       problemdetail.Acceptance,
				globalAcceptance: problemdetail.GlobalAcceptance,
			})
		}*/
		problemlist = problem
		code = vo.UnknownError

		return
	}
	code = vo.UserHasExisted
	log.Println("CreateMember:UserExisted")
	return
}

func (ctl ProblemController) ProblemDetail(c *gin.Context) {
	//TODO implement me
	panic("implement me")
	var req vo.ProblemDetailRequest
	var problem, p model.Problem
	var descriptionModel, d model.Description
	code := vo.OK
	var problemID = 0
	var description = ""
	var status = ""
	var difficulty = ""
	var acceptance = ""
	var globalAcceptance = ""
	defer func() {
		resp := vo.ProblemDetailResponse{
			Code:              code,
			ProblemID:         problemID,
			Description:       description,
			Status:            status,
			Difficulty:        difficulty,
			Acceptance:        acceptance,
			Global_Acceptance: globalAcceptance,
		}
		c.JSON(http.StatusOK, resp)
		utils.LogReqRespBody(req, resp, "ReturnProblemDescription")
	}()

	//descriptionModel = model.Description{ProblemID: req.ProblemID}
	problem = model.Problem{ProblemID: problemID, Status: status, Difficulty: difficulty, Acceptance: acceptance, GlobalAcceptance: globalAcceptance}
	if err := ctl.DB.Where("problem_id =?", req.ProblemID).Take(&p).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctl.DB.First(&problem, req.ProblemID)
			log.Println("Successfully find the problem detail information" + strconv.Itoa(problem.ProblemID))
			problemID = problem.ProblemID
			status = problem.Status
			difficulty = problem.Difficulty
			acceptance = problem.Acceptance
			globalAcceptance = problem.GlobalAcceptance
			return
		} else {
			code = vo.UnknownError
			log.Println("Problem Information :Unknown-error while finding Problem information")
			return
		}
	}
	//ctl.DB.Find(&problem)
	if err := ctl.DB.Where("problem_id =?", req.ProblemID).Take(&d).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctl.DB.First(&descriptionModel)
			log.Println("Successfully find the problem detail information" + string(descriptionModel.Description))
			description = string(descriptionModel.Description)
			return
		} else {
			code = vo.UnknownError
			log.Println("Problem Information :Unknown-error while finding Problem information")
			return
		}
	}
}
