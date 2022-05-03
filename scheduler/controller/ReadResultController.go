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
	list, err := service.ListRunningJob()
	if err != nil {
		logger.Errorln("List Job error")
		logger.Errorln(err)
		return
	}
	// logger.Infoln(strconv.Itoa(len(list.Items)) + " jobs are running.")
	logger.Infoln(list)

	// for _, job := range list.Items {
	// 	logger.Infoln(job.Status.Phase)
	// }
}
