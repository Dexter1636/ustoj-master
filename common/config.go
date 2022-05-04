package common

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func ReadConfig(configFile string) {
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf(err.Error()))
	}
	log.Printf("==== using config: %s ====", configFile)
}
