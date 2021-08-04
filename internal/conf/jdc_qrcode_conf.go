package conf

import "github.com/spf13/viper"

func GetJDQrcode() string {
	return viper.GetString("service.jdcHost")
}
