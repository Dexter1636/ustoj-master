package model

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Environment string     `yaml:"environment"`
	Server      Server     `yaml:"server"`
	Datasource  Datasource `yaml:"datasource"`
	DataPath    DataPath   `yaml:"dataPath"`
	Const       Const      `yaml:"const"`
	Logger      Logger     `yaml:"logger"`
	Scheduler   Scheduler  `yaml:"scheduler"`
	Kubernetes  Kubernetes `yaml:"kubernetes"`
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

type DataPath struct {
	SubmitPath string `yaml:"submitPath"`
}

type Const struct {
	Delimiter string `yaml:"delimiter"`
}

type Logger struct {
	Info      string `yaml:"info"`
	Level     string `yaml:"level"`
	WriteFile bool   `yaml:"writeFile"`
}

type Scheduler struct {
	DispatchInterval   int    `yaml:"dispatchInterval"`
	DispatchNum        int    `yaml:"dispatchNum"`
	ReadResultInterval int    `yaml:"readResultInterval"`
	JobPvcName         string `yaml:"jobPvcName"`
}

type Kubernetes struct {
	MasterUrl    string `yaml:"masrerUrl"`
	MasterConfig string `yaml:"masterConfig"`
}

var Cfg Config

func InitConfig() {
	if err := viper.Unmarshal(&Cfg); err != nil {
		panic(fmt.Errorf(err.Error()))
	}
	fmt.Println(Cfg)
}

func GetConfig() Config {
	return Cfg
}
