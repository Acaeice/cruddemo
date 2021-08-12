package etcode

import (
	"github.com/gin-gonic/gin"
	"github.com/wechatapi/cruddemo/pkg"
)

type qrCodeRepoI interface {
	GetQrcode(c *gin.Context) (*pkg.Qrcode, error)
	GetjQuery(c *gin.Context) (*pkg.Query, error)
	CheckLogin(token, cookie, okl_token string) string
}
