package service

import (
	"os"
	"strconv"
	"strings"
	"ustoj-master/common"
	"ustoj-master/dto"
	"ustoj-master/model"
	config "ustoj-master/scheduler/model"
)

func GetNWaitingSubmissions(n int, submissionDtoList *[]dto.SubmissionDto) {
	common.DB.Model(model.Submission{}).Where("status = ?", "submitted").Order("submission_time").Limit(n).Find(&submissionDtoList)
}

func UpdateSubmissionToAccepted(submission dto.SubmissionDto) {
	common.DB.Model(&model.Submission{}).Where("submission_id = ?", submission.SubmissionID).Update("status", "accepted")
}

func UpdateSubmissionToRunning(submission dto.SubmissionDto) {
	common.DB.Model(&model.Submission{}).Where("submission_id = ?", submission.SubmissionID).Update("status", "running")
}

func UpdateSubmissionToRuntimeError(submission dto.SubmissionDto) {
	common.DB.Model(&model.Submission{}).Where("submission_id = ?", submission.SubmissionID).Update("status", "runtimeError")
}

func UpdateSubmissionToWrongAnswer(submission dto.SubmissionDto) {
	common.DB.Model(&model.Submission{}).Where("submission_id = ?", submission.SubmissionID).Update("status", "wrongAnswer")
}

func UpdateSubmissionToInternalError(submission dto.SubmissionDto) {
	common.DB.Model(&model.Submission{}).Where("submission_id = ?", submission.SubmissionID).Update("status", "internalError")
}

func UpdateSubmissionsToPending(submissionList *[]dto.SubmissionDto) {
	for _, sub := range *submissionList {
		UpdateSubmissionToRunning(sub)
	}
}

func GetCaseListByProblemId(problemId int, caseList *[]string) {
	common.DB.Table("test_case").Select("case").Where("problem_id = ?", problemId).Order("case_id").Scan(&caseList)
}

func GetExpectedListByProblemId(problemId int, exceptedList *[]string) {
	common.DB.Table("test_case").Select("expected").Where("problem_id = ?", problemId).Order("case_id").Scan(&exceptedList)
}

func WriteCodeToFile(code string, filePath string) error {
	if IsExists(filePath) {
		if err := os.Remove(filePath); err != nil {
			logger.Errorln(err.Error())
		}
	}
	f, err := os.Create(filePath)
	defer f.Close()
	if err != nil {
		logger.Errorln(err.Error())
		return err
	} else {
		_, err = f.Write([]byte(code))
		logger.Errorln(err.Error())
		return err
	}
}

func IsExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func CheckResult(submissionId int, problemId int) (bool, error) {
	cfg := config.GetConfig()
	// read output
	buf, err := os.ReadFile(cfg.DataPath.SubmitPath + strconv.Itoa(submissionId) + "/output")
	if err != nil {
		return false, err
	}
	outputStr := string(buf)
	outputList := strings.Split(outputStr, cfg.Const.Delimiter)
	// read expected
	exceptedList := make([]string, 0, 8)
	GetExpectedListByProblemId(problemId, &exceptedList)
	// compare
	if len(outputList) != len(exceptedList) {
		return false, nil
	}
	for i := 0; i < len(outputList); i++ {
		if strings.TrimSpace(outputList[i]) != strings.TrimSpace(exceptedList[i]) {
			return false, nil
		}
	}
	return true, nil
}
