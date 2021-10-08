package cruddemo

import (
	"github.com/gin-gonic/gin"
	"github.com/wechatapi/cruddemo/internal/etcode"
)

func getQrcode(c *gin.Context) {
	qrcode, err := etcode.QrcodeRepo.GetQrcode(c)
	if err != nil {
		fail(c, err)
		return
	}
	ok(c, resp{
		"qrcode": qrcode,
	})
}

func getjQuery(c *gin.Context) {
	jquery, err := etcode.QrcodeRepo.GetjQuery(c)
	if err != nil {
		fail(c, err)
		return
	}
	ok(c, resp{
		"jquery": jquery,
	})
}
