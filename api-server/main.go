package main

import (
	"os"
	"ustoj-master/api-server/model"
	"ustoj-master/common"
	appConfig "ustoj-master/scheduler/model"

	"github.com/spf13/viper"
)

func main() {
	common.ReadConfig(os.Args[1])
	model.InitConfig()
	common.InitLogger(appConfig.Cfg.Logger.WriteFile)
	common.InitDb(appConfig.Cfg.Logger.Level)
	r := RegisterRouter()
	port := viper.GetString("server.port")
	r.Run("0.0.0.0:" + port)
}
