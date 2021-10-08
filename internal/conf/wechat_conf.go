package conf

import "github.com/spf13/viper"

// Database 数据库配置
type Wechat struct {
	AppID     string
	AppSecret string
}

func GetWechat() *Wechat {
	return &Wechat{
		AppID:     viper.GetString("wechatAPP.appID"),
		AppSecret: viper.GetString("wechatAPP.appSecret"),
	}
}
