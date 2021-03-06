package util

import (
	"log"

	"github.com/beego/beego/v2/client/httplib"
)

var ua = "Mozilla/5.0 (iPhone; CPU iPhone OS 13_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 SP-engine/2.14.0 main%2F1.0 baiduboxapp/11.18.0.16 (Baidu; P2 13.3.1) NABar/0.0"

func UserAgent() {
	log.Print("更新User-Agent")
	var err error
	ua, err = httplib.Get("https://cdn.jsdelivr.net/gh/cdle/xdd@main/ua.txt").String()
	if err != nil {
		log.Print("更新User-Agent失败")
	}
}

func GetUserAgent() string {
	return ua
}
