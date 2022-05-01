package controller

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"ustoj-master/dto"
	"ustoj-master/scheduler/model"
	"ustoj-master/service"

	"ustoj-master/common"
)

var logger = common.LogInstance()
var cfg = model.Cfg

func RunDispatch(done func()) {
	defer done()

	c := cron.New(cron.WithSeconds())

	spec := "*/2 * * * * ?"
	c.AddFunc(spec, func() {
		logger.Infoln("cron RunDispatch")
		// acquire n submissions
		submissionList := make([]dto.SubmissionDto, 0, cfg.Scheduler.SubmissionNum)
		service.GetNWaitingSubmissions(cfg.Scheduler.SubmissionNum, &submissionList)
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
