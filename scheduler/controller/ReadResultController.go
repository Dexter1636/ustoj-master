package controller

import (
	"github.com/robfig/cron/v3"
)

func RunReadResult(done func()) {
	defer done()

	c := cron.New(cron.WithSeconds())

	spec := "*/2 * * * * ?"
	c.AddFunc(spec, func() {
		logger.Infoln("cron RunReadResult")
	})

	c.Start()
	select {} // block
}
