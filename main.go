package main

import (
	"github.com/spf13/viper"
	"ustoj-master/common"
)

func main() {
	common.InitConfig("application")
	common.InitLogger()
	//common.InitDb()
	r := RegisterRouter()
	port := viper.GetString("server.port")
	r.Run("0.0.0.0:" + port)
}
