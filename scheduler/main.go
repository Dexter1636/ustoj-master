package main

import (
	"os"
	"ustoj-master/common"
	"ustoj-master/scheduler/model"
)

func main() {
	common.ReadConfig(os.Args[1])
	model.InitConfig()
	common.InitLogger()
	//common.InitDb()
}
