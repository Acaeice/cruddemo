package cruddemo

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
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
	store := memstore.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		// MaxAge:   240,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	})
	freecar := api.Group("api/v1")
	{
		freecar.Use(Cors())
		freecar.GET("/code", getCode)
		freecar.POST("/login", wechatQuickLogin)
		freecar.GET("", answer)
		freecar.Use(credential)
		freecar.Use(auth)
		freecar.GET("/user", getuserbyId)
		freecar.Use(sessions.Sessions("ET", store))
		freecar.GET("/jquery", getjQuery)
		freecar.GET("/qrcode", getQrcode)
	}
}
