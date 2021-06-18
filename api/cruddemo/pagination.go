package cruddemo

import (
	"github.com/gin-gonic/gin"
	"github.com/wechatapi/cruddemo/pkg"
	"github.com/wechatapi/cruddemo/util"
)

// defaultPage defaultPageSize 每页数据条数
const (
	defaultPage     = 1
	defaultPageSize = 10
)

// paginationByGet 通过当前页码新建一个Pagination
func paginationByGet(c *gin.Context) *pkg.Pagination {
	page := c.Query("page")
	pageSize := c.Query("pageSize")
	var p uint = defaultPage
	if page != "" {
		temp, err := util.ParseUint(page)
		if err == nil && temp >= 1 {
			p = temp
		}
	}
	var ps uint = defaultPageSize
	if pageSize != "" {
		temp, err := util.ParseUint(pageSize)
		if err == nil && temp >= 1 {
			ps = temp
		}
	}
	return &pkg.Pagination{
		Total:    0,
		PageSize: ps,
		Current:  p,
		Offset:   (p - 1) * ps,
	}
}
