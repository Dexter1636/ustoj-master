package controller

import (
	"ustoj-master/service"

	"github.com/robfig/cron/v3"
)

func RunReadResult(done func()) {
	defer done()

	c := cron.New(cron.WithSeconds())

	spec := "*/2 * * * * ?"
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
