package cruddemo

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wechatapi/cruddemo/internal/shared"
	"github.com/wechatapi/cruddemo/internal/user"
	"github.com/wechatapi/cruddemo/util"
)

// extractToken 提取token
func extractToken(c *gin.Context) string {
	headerToken := c.Request.Header.Get("token")
	if len(headerToken) > 0 {
		log.Printf("header token: %s", headerToken)
		return headerToken
	}
	paramToken := c.Query("token")
	if len(paramToken) > 0 {
		log.Printf("param token: %s", headerToken)
		return paramToken
	}
	return ""
}

// 根据token获取用户的登录信息
func credential(c *gin.Context) {
	token := extractToken(c)

	// 如果token的长度为0，说明用户未登录过
	if len(token) == 0 {
		c.Next()
		return
	}

	userID, err := util.Unsign(token)
	if err != nil {
		log.Print(err)
	}
	userProfile, err := user.UserRepo.GetByID(userID)
	if err != nil {
		c.Next()
		return
	}

	ctx := shared.WithUserProfile(c.Request.Context(), userProfile)
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}

// auth 需要登录拦截
func auth(c *gin.Context) {
	_, ok := shared.GetUserProfile(c.Request.Context())
	if !ok {
		unLogin(c)
		c.Abort()
	}

	c.Next()
}

// logUserVisit 记录用户活跃信息
func logUserVisit(c *gin.Context) {
	type ActivityMsg struct {
		UserID  int  `json:"userID" form:"userID"`
		MallID  uint `json:"mallID" form:"mallID"`
		StoreID uint `json:"storeID" form:"storeID"`
		GoodsID uint `json:"goodsID" form:"goodsID"`
	}
	p := ActivityMsg{}
	if err := c.ShouldBindQuery(&p); err != nil {
		log.Printf("error: 【full】 %+#v ", err)
		return
	}
	u, ok := shared.GetUserProfile(c.Request.Context())
	if ok {
		p.UserID = int(u.ID)
	} else {
		c.ClientIP()
		parts := strings.Split(c.ClientIP(), ".")
		result := strings.Join(parts, "")
		intRes, _ := strconv.Atoi(result)
		p.UserID = -intRes
	}

	// 发送使用消息
	// go func() {
	// 	prod := kafka.NewProducer()
	// 	if err := prod.Errors(); err != nil {
	// 		log.Printf("用户活跃信息数据投递失败, 请检查Kafka初始化\nerror:%+#v", err)
	// 		return
	// 	}
	// 	prod.Produce(kafka.TrackingUserVisit, fmt.Sprintf("%03v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100)), p, time.Now())
	// }()

	c.Next()
}

func CorsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // 这是允许访问所有域
		// c.Header().Set("Access-Control-Allow-Origin", "")           // 跨域请求是否需要带cookie信息 默认设置为true
		c.Header("Access-Control-Allow-Credentials", "true") // 跨域请求是否需要带cookie信息 默认设置为true
		//  header的类型
		c.Header("Access-Control-Allow-Origin", "*") // 这是允许访问所有域
		// c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control,XMLHttpRequest, X-Requested-With")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, custom-header, Cache-Control,XMLHttpRequest, X-Requested-With, token")
		//c.Header("Access-Control-Allow-Headers", "*")
		//服务器支持的所有跨域请求的方法
		// c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		// c.Header("Access-Control-Max-Age", "21600") //可以缓存预检请求结果的时间（以秒为单位）
		// c.Set("content-type", "application/json")   // 设置返回格式是json
		if c.Request.Method == "OPTIONS" {
			// 	// c.AbortWithStatus(204)
			// 	c.AbortWithStatus(http.StatusNoContent)
			// 	return
			// 	放行所有OPTIONS方法，本项目直接返回204
			c.JSON(200, "Options Request!")
		}

		c.Next()
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("Cors() \n")
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}
