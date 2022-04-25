package main

import (
	"os"
	"sync"
	"ustoj-master/common"
	"ustoj-master/scheduler/controller"
	"ustoj-master/scheduler/model"
)

func main() {
	common.ReadConfig(os.Args[1])
	model.InitConfig()
	common.InitLogger(model.Cfg.Logger.WriteFile)
	common.InitDb(model.Cfg.Logger.Level)

	var wg sync.WaitGroup
	wg.Add(1)
	go controller.RunDispatch(wg.Done)
	go controller.RunReadResult(wg.Done)
	wg.Wait()
}
