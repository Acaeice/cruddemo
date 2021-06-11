package conf

import "github.com/spf13/viper"

// LogConf 日志配置
type LogConf struct {
	Engine        string // 日志引擎
	EnableConsole bool   // 是否输出控制台日志
	Level         string // 文件日志级别
	FileLocation  string // 日志文件路径
	EnableFile    bool   // 是否在文件记录日志
	AppendCaller  bool   // 是否打印代码位置
}

// GetLogConfig 获取gin配置
func GetLogConfig() *LogConf {
	return &LogConf{
		Engine:        viper.GetString("log.engine"),
		EnableConsole: viper.GetBool("log.enableConsole"),
		Level:         viper.GetString("log.level"),
		FileLocation:  viper.GetString("log.fileLocation"),
		EnableFile:    viper.GetBool("log.enableFile"),
		AppendCaller:  viper.GetBool("log.appendCaller"),
	}
}
