package model

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Environment string     `yaml:"environment"`
	Server      Server     `yaml:"server"`
	Datasource  Datasource `yaml:"datasource"`
	Logger      Logger     `yaml:"logger"`
}

type Server struct {
	Port string `yaml:"port"`
}

type Datasource struct {
	DriverName string `yaml:"driverName"`
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	Database   string `yaml:"database"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	Charset    string `yaml:"charset"`
}

type Logger struct {
	Info string `yaml:"info"`
}

var Cfg Config

func InitConfig() {
	if err := viper.Unmarshal(&Cfg); err != nil {
		panic(fmt.Errorf(err.Error()))
	}
}
