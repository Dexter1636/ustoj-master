package main

import (
	"flag"
	"sync"
	"ustoj-master/common"
	commonModel "ustoj-master/model"
	"ustoj-master/scheduler/controller"
	"ustoj-master/scheduler/model"
	appConfig "ustoj-master/scheduler/model"
	cluster "ustoj-master/service"
)

func main() {
	var configPath = flag.String("c", "/etc/ustoj/master-ticker/config.yaml", "the file path to config file")
	flag.Parse()

	common.ReadConfig(*configPath)
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
	common.CreateTableIfNotExists(commonModel.Problem{})
	common.CreateTableIfNotExists(commonModel.Description{})
	common.CreateTableIfNotExists(commonModel.TestCase{})
}
