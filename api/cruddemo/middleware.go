package cruddemo

import (
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

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
