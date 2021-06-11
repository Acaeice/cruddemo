package cruddemo

import (
	"github.com/gin-gonic/gin"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
)

func initRouter(router *gin.Engine) {
	// debug模式下打印请求日志
	// logConf := conf.GetLogConfig()
	// if logConf.Level == "debug" {
	// 	router.Use(logger.LogRequest())
	// }

	// api文档，地址http://localhost:8080/swagger/
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("")
	initFreeCarRouter(api)
}

func initFreeCarRouter(api *gin.RouterGroup) {
	freecar := api.Group("api/v1")
	{
		freecar.GET("/code", getCode)
		freecar.POST("/login", wechatQuickLogin)
		freecar.GET("/user", getuserbyId)
	}
}
