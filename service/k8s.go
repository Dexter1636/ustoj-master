package service

import (
	"ustoj-master/common"
	"ustoj-master/model"
)

var logger = common.LogInstance()

var c model.Cluster

func InitCluster(masterUrl string, masterConfigPath string) {
	c.InitKube(masterUrl, masterConfigPath)
}
