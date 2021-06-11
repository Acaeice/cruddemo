package conf

import "github.com/spf13/viper"

// Database 数据库配置
type Database struct {
	Dsn     string
	Debug   bool
}

// GetDatabase 获取数据库配置
func GetDatabase() *Database {
	return &Database{
		Dsn:     viper.GetString("database.dsn"),
		Debug:   viper.GetBool("database.debug"),
	}
}
