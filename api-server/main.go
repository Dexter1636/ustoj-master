package main

import (
	"github.com/spf13/viper"
	"os"
	"ustoj-master/common"
)

func main() {
	common.InitConfig(os.Args[1])
	common.InitLogger()
	//common.InitDb()
	r := RegisterRouter()
	port := viper.GetString("server.port")
	r.Run("0.0.0.0:" + port)
}
