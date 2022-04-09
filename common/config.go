package common

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

func InitConfig(env string) {
	workDir, _ := os.Getwd()
	viper.SetConfigName(env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf(err.Error()))
	}
	log.Printf("========using %s config========", env)
}
