package service

import (
	"ustoj-master/common"
	"ustoj-master/dto"
	"ustoj-master/model"
)

func GetNWaitingSubmissions(n int, submissionDtoList *[]dto.SubmissionDto) {
	common.DB.Model(model.Submission{}).Where("status = ?", "submitted").Order("submission_time").Limit(n).Find(&submissionDtoList)
}

func UpdateSubmissionsToPending(submissionList *[]dto.SubmissionDto) {
	for _, dto := range *submissionList {
		common.DB.Model(&model.Submission{}).Where("submission_id = ?", dto.SubmissionID).Update("status", "pending")
	}
}

func GetCaseListByProblemId(problemId int64, caseList *[]string) {
	common.DB.Table("case").Select("case").Where("problem_id = ?", problemId).Scan(&caseList)
}

func GetResultListByProblemId(problemId int64, resultList *[]string) {
	common.DB.Table("case").Select("result").Where("problem_id = ?", problemId).Scan(&resultList)
}
