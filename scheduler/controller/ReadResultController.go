package controller

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func RunReadResult(done func()) {
	defer done()

	c := cron.New(cron.WithSeconds())

	spec := "*/2 * * * * ?"
	c.AddFunc(spec, func() {
		fmt.Println("cron RunReadResult")
	})

	c.Start()
	select {} // block
}
