package service

import (
	"context"
	"errors"
	"log"
	"strconv"
	"ustoj-master/common"
	"ustoj-master/model"
	"ustoj-master/vo"

	"gorm.io/gorm"
)

type DBService interface {
	CreateUser(user *model.User) vo.ErrNo
	Login(user *model.User) string
	//GetProblemList(map[string]interface{}) ([]*model.Problem, error)
	GetProblemList(problem []model.Problem) []model.Problem
	ProblemDetail(problemID int) model.Problem
	ProblemDescription(problemID int) model.Description
	Submission(submission *model.Submission)
	ResultList(submitted *model.Submission) model.Submission
}
type DBConnect struct {
	DB  *gorm.DB
	Ctx context.Context
}

//func NewDBConnect(DB *gorm.DB) DBService { // Similar to the interface of service
func NewDBConnect() DBService { // Similar to the interface of service
	return &DBConnect{DB: common.GetDB(), Ctx: common.GetCtx()}
}

/*func NewDBService() DBService {
	 retrun
}*/

func (db *DBConnect) CreateUser(user *model.User) vo.ErrNo {
	var u model.User
	if err := db.DB.Where("username = ?", user.Username).Take(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			db.DB.Create(&user)
			log.Println("CreateMember:Successfully create, username:" + user.Username)
			return vo.OK
		} else {
			log.Println("CreateMember:Unknown-error while creating")
			return vo.UnknownError
		}
	}
	return vo.UnknownError
}

func (db *DBConnect) Login(user *model.User) string {
	var u model.User
	if err := db.DB.Where("username = ?", user.Username).Where("password =?", user.Password).Take(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Not User record")
			return "UnknownError"
		} else {
			return "UnknownError"
		}
	}
	return u.Username
}
func (db *DBConnect) GetProblemList(problem []model.Problem) []model.Problem {
	var problemlist []model.Problem
	if result := db.DB.Find(&problem); result.Error != nil {
		log.Println("Error occured during get all problem information! ")
		return problemlist
	} else {

		log.Println("The lenght of all problem :" + strconv.FormatInt(result.RowsAffected, 10))

		problemlist = problem

		return problemlist
	}
}
func (db *DBConnect) ProblemDetail(problemID int) model.Problem {
	var problem, p model.Problem
	if err := db.DB.Where("problem_id =?", problemID).Take(&p).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			db.DB.First(&problem, problemID)
			log.Println("Successfully find the problem detail information" + strconv.Itoa(problem.ProblemID))

			return problem
		} else {

			return problem
		}
	}
	return problem
}
func (db *DBConnect) ProblemDescription(problemID int) model.Description {
	var descriptionModel, d model.Description
	if err := db.DB.Where("problem_id =?", problemID).Take(&d).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			db.DB.First(&descriptionModel)
			log.Println("Successfully find the problem detail information" + string(descriptionModel.Description))

			return descriptionModel
		} else {

			return descriptionModel
		}
	}
	return descriptionModel
}
func (db *DBConnect) Submission(submission *model.Submission) {
	if err := db.DB.Create(&submission).Error; err != nil {
		log.Println("Submission Error:" + err.Error())
	}

}
func (db *DBConnect) ResultList(submitted *model.Submission) model.Submission {
	var submission, s model.Submission
	if err := db.DB.Where("problem_id =?", submitted.ProblemID).Where("username = ?", submitted.Username).Take(&s).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			db.DB.First(&submission)
			log.Println("Successfully find the problem result" + strconv.Itoa(submission.ProblemID))

			return submission
		} else {

			return submission
		}
	}
	return submission

}
