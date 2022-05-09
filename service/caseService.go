package service

import (
	"io/ioutil"
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
	common.DB.Model(&model.Submission{}).Where("submission_id = ?", submission.SubmissionID).Update("status", "runtime error")
}

func UpdateSubmissionToWrongAnswer(submission dto.SubmissionDto) {
	common.DB.Model(&model.Submission{}).Where("submission_id = ?", submission.SubmissionID).Update("status", "wrong answer")
}

func UpdateSubmissionToInternalError(submission dto.SubmissionDto) {
	common.DB.Model(&model.Submission{}).Where("submission_id = ?", submission.SubmissionID).Update("status", "internal error")
}

func UpdateSubmissionsToPending(submissionList *[]dto.SubmissionDto) {
	for _, sub := range *submissionList {
		UpdateSubmissionToRunning(sub)
	}
}

func GetCaseListByProblemId(problemId int, caseList *[]string) {
	common.DB.Table("test_case").Select("`case`").Where("problem_id = ?", problemId).Order("case_id").Scan(&caseList)
}

func GetExpectedListByProblemId(problemId int, exceptedList *[]string) {
	common.DB.Table("test_case").Select("`expected`").Where("problem_id = ?", problemId).Order("case_id").Scan(&exceptedList)
}

// WriteCodeToFile writes the code text to related file.
// It checks whether the related file exists.
// If yes, it will remove it and create a new one. If not, it will create the file directly.
func WriteCodeToFile(code string, dirPath string) error {
	if IsExists(dirPath) {
		if err := os.RemoveAll(dirPath); err != nil {
			logger.Errorln("error when remove file" + err.Error())
		}
	}
	logger.Infoln("create file")
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}
	f, err := os.Create(dirPath + "/code")
	if err != nil {
		return err
	}
	defer f.Close()
	if err != nil {
		return err
	} else {
		_, err = f.Write([]byte(code))
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
	buf, err := ioutil.ReadFile(cfg.DataPath.SubmitPath + strconv.Itoa(submissionId) + "/output/output.txt")
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
