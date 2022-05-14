package main

import (
	"flag"
	"ustoj-master/api-server/model"
	"ustoj-master/common"
	commonModel "ustoj-master/model"
	appConfig "ustoj-master/scheduler/model"

	"github.com/spf13/viper"
)

func main() {
	var configPath = flag.String("c", "/etc/ustoj/master-apiserver/config.yaml", "the file path to config file")
	flag.Parse()

	common.ReadConfig(*configPath)
	model.InitConfig()
	common.InitLogger(appConfig.Cfg.Logger.WriteFile)
	common.InitDb(appConfig.Cfg.Logger.Level)
	InitTable()
	r := RegisterRouter()
	port := viper.GetString("server.port")
	r.Run("0.0.0.0:" + port)
}

func InitTable() {
	common.CreateTableIfNotExists(commonModel.User{})
	common.CreateTableIfNotExists(commonModel.Problem{})
	common.CreateTableIfNotExists(commonModel.Description{})
	common.CreateTableIfNotExists(commonModel.Submission{})
	common.CreateTableIfNotExists(commonModel.TestCase{})
}
