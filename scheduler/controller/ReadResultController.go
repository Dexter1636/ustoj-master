package controller

import (
	"fmt"
	"ustoj-master/scheduler/model"
	"ustoj-master/service"

	"github.com/robfig/cron/v3"
)

func RunReadResult(done func()) {
	defer done()
	cfg := model.GetConfig()

	c := cron.New(cron.WithSeconds())

	spec := fmt.Sprintf("*/%d * * * * ?", cfg.Scheduler.ReadResultInterval)
	c.AddFunc(spec, MainJob)

	c.Start()
	select {} // block
}

func MainJob() {
	list, err := service.ListJobById([]int{1})
	if err != nil {
		logger.Errorln("List Job error")
		logger.Errorln(err)
		return
	}
	logger.Infoln("=== job list:")
	logger.Infoln(list)

}
