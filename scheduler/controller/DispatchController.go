package controller

import (
	"github.com/robfig/cron/v3"

	"ustoj-master/common"
)

var logger = common.LogInstance()

func RunDispatch(done func()) {
	defer done()

	c := cron.New(cron.WithSeconds())

	spec := "*/2 * * * * ?"
	c.AddFunc(spec, func() {
		logger.Infoln("cron RunDispatch")
		// TODO: acquire n submissions

		// TODO: acquire related info and call k8s service to run the jobs
		// _code, case, _lang

	})

	c.Start()
	select {} // block
}
