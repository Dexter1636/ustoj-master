package service

import (
	"context"
	"ustoj-master/model"

	"gorm.io/gorm"
)

type DBService interface {
	Login(Username string, Password string) string
	GetProblemList(map[string]interface{}) ([]*model.Problem, error)
	ProblemDetail(map[string]interface{}) (*model.Problem, error)
	/*...........Other still missing..........*/
}
type DBConnect struct {
	DB  *gorm.DB
	Ctx context.Context
}

/*func NewDBConnect(db *gorm.DB) DBService { // Similar to the interface of service
	return &DBConnect{DB: common.GetDB(), Ctx: common.GetCtx()}
}*/

/*func NewDBService() DBService {
	 retrun
}*/

func Login(Username string, Password string) string {
	test := ""
	/*if err := DBConnect.DB.Where("user_name = ?", Username).Where("password =?", Password).Take(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			loginobject.Data.UserID = u.Username
		} else {
			code = vo.UnknownError
			log.Println("Not User record")
			return
		}
	}*/
	return test
}
