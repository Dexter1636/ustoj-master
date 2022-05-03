package controller

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"ustoj-master/common"
	"ustoj-master/dto"
	"ustoj-master/scheduler/model"
	"ustoj-master/service"
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
		submissionList := make([]dto.SubmissionDto, 0, cfg.Scheduler.DispatchNum)
		service.GetNWaitingSubmissions(cfg.Scheduler.DispatchNum, &submissionList)
		// update acquired submissions to status pending
		service.UpdateSubmissionsToPending(&submissionList)
		// acquire related info and call k8s service to run the jobs
		// _code, caseList, _lang
		for _, sub := range submissionList {
			code := sub.Code
			caseList := make([]string, 0, 8)
			lang := sub.Language
			fmt.Println(code, caseList, lang)
			// TODO: call k8s service to run the jobs
		}
	})

	c.Start()
	select {} // block
}
