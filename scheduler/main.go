package main

import (
	"os"
	"ustoj-master/common"
)

func main() {
	common.InitConfig(os.Args[1])
	common.InitLogger()
	//common.InitDb()
}
