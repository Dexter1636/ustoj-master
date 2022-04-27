package service

import (
	"ustoj-master/common"
)

var DB = common.GetDB()

func GetCaseListByProblemId(problemId int64, caseList *[]string) {
	DB.Table("case").Select("case").Where("problem_id = ?", problemId).Scan(&caseList)
}

func GetResultListByProblemId(problemId int64, resultList *[]string) {
	DB.Table("case").Select("result").Where("problem_id = ?", problemId).Scan(&resultList)
}
