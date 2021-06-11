package conf

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const (
	linuxCfg   = "/var/config/config.yaml"
	devCfg     = "./config.yaml"
	projectCfg = "/src/code.meikeland.com/wanghejun/cruddemo/config.yaml"
)

// Init 读取配置文件
// 会依次到 linux默认路径 > windows默认路径 > local 寻找配置文件，全部没找到时候panic
func Init() {
	cfgFile := linuxCfg
	if _, err := os.Stat(cfgFile); err == nil {
		initViper(cfgFile)
		return
	}

	cfgFile = devCfg
	if _, err := os.Stat(cfgFile); err == nil {
		initViper(cfgFile)
		return
	}

	cfgFile = os.Getenv("GOPATH") + projectCfg
	if _, err := os.Stat(cfgFile); err != nil {
		panic("Find config error: " + err.Error())
	}
	initViper(cfgFile)
}

func initViper(file string) {
	viper.SetConfigFile(file)
	if err := viper.ReadInConfig(); err != nil {
		panic("Read config error: " + err.Error())
	}

	fmt.Println("Using config file:", viper.ConfigFileUsed())
}
