package main

import (
	"github.com/spf13/viper"
	"os"
	"ustoj-master/api-server/model"
	"ustoj-master/common"
)

func main() {
	common.ReadConfig(os.Args[1])
	model.InitConfig()
	common.InitLogger()
	common.InitDb()
	r := RegisterRouter()
	port := viper.GetString("server.port")
	r.Run("0.0.0.0:" + port)
}
