package common

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func ReadConfig(configFile string) {
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf(err.Error()))
	}
	log.Printf("==== using config: %s ====", configFile)
}
