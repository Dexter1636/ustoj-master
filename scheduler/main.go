package main

import (
	"os"
	"sync"
	"ustoj-master/common"
	commonModel "ustoj-master/model"
	"ustoj-master/scheduler/controller"
	"ustoj-master/scheduler/model"
	appConfig "ustoj-master/scheduler/model"
	cluster "ustoj-master/service"
)

func main() {
	common.ReadConfig(os.Args[1])
	model.InitConfig()
	common.InitLogger(appConfig.Cfg.Logger.WriteFile)
	common.InitDb(appConfig.Cfg.Logger.Level)
	InitTable()
	cluster.InitCluster(appConfig.Cfg.Kubernetes.MasterUrl, appConfig.Cfg.Kubernetes.MasterConfig)

	var wg sync.WaitGroup
	wg.Add(1)
	go controller.RunDispatch(wg.Done)
	go controller.RunReadResult(wg.Done)
	wg.Wait()
}

func InitTable() {
	common.CreateTableIfNotExists(commonModel.Submission{})
}
