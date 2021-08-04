package cruddemo

import (
	"log"

	"github.com/Acaeice/jdc/sdk/jdcsdk"
	"github.com/gin-gonic/gin"
	"github.com/meikeland/logger"
	"github.com/wechatapi/cruddemo/internal/conf"
	"github.com/wechatapi/cruddemo/internal/sql"
	"github.com/wechatapi/cruddemo/internal/user"
)

// Init 初始化服务
func Init() {

	conf.Init() // 配置的初始化
	sql.Init()  // 数据库初始化

	initLogger()

	// 其他需要初始化的sdk和internal包在这个位置完成
	user.Init()
	jdcsdk.Init(conf.GetJDQrcode())
	// 路由初始化
	gin.SetMode(conf.GetGin().Mode)
	r := gin.Default()
	initRouter(r)
	log.Fatal(r.Run(":8081"))
}

// initLogger 初始化日志引擎
func initLogger() {
	logConf := conf.GetLogConfig()

	config := logger.Config{
		EnableConsole: logConf.EnableConsole,
		Level:         logConf.Level,
		EnableFile:    logConf.EnableFile,
		FileLocation:  logConf.FileLocation,
		AppendCaller:  logConf.AppendCaller,
	}

	err := logger.New(config)
	if err != nil {
		log.Fatalf("Could not instantiate log %s", err.Error())
	}
}
