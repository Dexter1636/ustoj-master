package controller

import (
	"fmt"
	"ustoj-master/model"
	config "ustoj-master/scheduler/model"
	"ustoj-master/service"

	"github.com/robfig/cron/v3"
)

func RunReadResult(done func()) {
	defer done()
	cfg := config.GetConfig()

	c := cron.New(cron.WithSeconds())

	spec := fmt.Sprintf("*/%d * * * * ?", cfg.Scheduler.ReadResultInterval)
	c.AddFunc(spec, ReadResultJob)

	c.Start()
	select {} // block
}

func ReadResultJob() {
	subDtoList, err := service.ListJob()
	if err != nil {
		logger.Errorln("List Job error")
		logger.Errorln(err)
		return
	}
	logger.Infoln("=== job subDtoList:")
	logger.Infoln(subDtoList)
	for _, subDto := range subDtoList {
		switch subDto.Status {
		case model.JobSuccess:
			isRightAnswer, err := service.CheckResult(subDto.SubmissionID)
			if err != nil {
				service.UpdateSubmissionToInternalError(subDto)
			} else if isRightAnswer {
				service.UpdateSubmissionToAccepted(subDto)
			} else {
				service.UpdateSubmissionToWrongAnswer(subDto)
			}
		case model.JobFailed:
			service.UpdateSubmissionToRuntimeError(subDto)
		case model.JobUnknown:
			service.UpdateSubmissionToInternalError(subDto)
		}
	}
}
