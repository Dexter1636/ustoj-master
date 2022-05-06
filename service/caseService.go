package service

import (
	"os"
	"ustoj-master/common"
	"ustoj-master/dto"
	"ustoj-master/model"
)

func GetNWaitingSubmissions(n int, submissionDtoList *[]dto.SubmissionDto) {
	common.DB.Model(model.Submission{}).Where("status = ?", "submitted").Order("submission_time").Limit(n).Find(&submissionDtoList)
}

func UpdateSubmissionToPending(submission dto.SubmissionDto) {
	common.DB.Model(&model.Submission{}).Where("submission_id = ?", submission.SubmissionID).Update("status", "pending")
}

func UpdateSubmissionToInternalError(submission dto.SubmissionDto) {
	common.DB.Model(&model.Submission{}).Where("submission_id = ?", submission.SubmissionID).Update("status", "internalError")
}

func UpdateSubmissionsToPending(submissionList *[]dto.SubmissionDto) {
	for _, sub := range *submissionList {
		UpdateSubmissionToPending(sub)
	}
}

func GetCaseListByProblemId(problemId int64, caseList *[]string) {
	common.DB.Table("case").Select("case").Where("problem_id = ?", problemId).Scan(&caseList)
}

func GetResultListByProblemId(problemId int64, resultList *[]string) {
	common.DB.Table("case").Select("result").Where("problem_id = ?", problemId).Scan(&resultList)
}

func WriteCodeToFile(code string, filePath string) error {
	var f *os.File
	var err error
	if IsExists(filePath) {
		f, err = os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0600)
	} else {
		f, err = os.Create(filePath)
	}
	defer f.Close()
	if err != nil {
		logger.Error(err.Error())
		return err
	} else {
		_, err = f.Write([]byte(code))
		logger.Error(err.Error())
		return err
	}
}

func IsExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
