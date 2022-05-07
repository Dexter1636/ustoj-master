package controller

import (
	"fmt"
	"strconv"
	"ustoj-master/common"
	"ustoj-master/dto"
	"ustoj-master/scheduler/model"
	"ustoj-master/service"

	"github.com/robfig/cron/v3"
)

var logger = common.LogInstance()

func RunDispatch(done func()) {
	defer done()
	cfg := model.GetConfig()

	c := cron.New(cron.WithSeconds())
	spec := fmt.Sprintf("*/%d * * * * ?", cfg.Scheduler.DispatchInterval)
	c.AddFunc(spec, func() {
		logger.Infoln("cron RunDispatch")
		// acquire n submissions
		submissionDtoList := make([]dto.SubmissionDto, 0, cfg.Scheduler.DispatchNum)
		service.GetNWaitingSubmissions(cfg.Scheduler.DispatchNum, &submissionDtoList)
		// acquire related info and call k8s service to run the jobs
		// _code, caseList, _lang
		for _, subDto := range submissionDtoList {
			subId := subDto.SubmissionID
			code := subDto.Code
			caseList := make([]string, 0, 8)
			service.GetCaseListByProblemId(subDto.ProblemID, &caseList)
			lang := subDto.Language
			fmt.Println(code, caseList, lang)
			// write code snippet to file system
			if err := service.WriteCodeToFile(code, cfg.DataPath.SubmitPath+strconv.Itoa(subId)+"/code"); err != nil {
				service.UpdateSubmissionToInternalError(subDto)
			}
			// call k8s service to run the jobs
			if err := service.CreateJob(subId, caseList, lang); err != nil {
				service.UpdateSubmissionToInternalError(subDto)
			} else {
				// update acquired submissions to status pending
				service.UpdateSubmissionToPending(subDto)
			}
		}
	})

	c.Start()
	select {} // block
}
