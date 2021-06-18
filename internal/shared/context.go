package shared

import (
	"context"

	"github.com/wechatapi/cruddemo/pkg"
)

type key string

const userKey key = "user"

// WithUserProfile 往已有的context中注入管理员信息
func WithUserProfile(ctx context.Context, admin *pkg.User) context.Context {
	return context.WithValue(ctx, userKey, admin)
}

// GetUserProfile 从context中提取管理员
func GetUserProfile(ctx context.Context) (*pkg.User, bool) {
	admin, ok := ctx.Value(userKey).(*pkg.User)
	return admin, ok
}
