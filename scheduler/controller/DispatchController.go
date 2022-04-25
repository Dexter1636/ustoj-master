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
	})

	c.Start()
	select {} // block
}
