package shared

import (
	"context"

	"github.com/wechatapi/cruddemo/pkg"
)

type key string

const userKey key = "user"
const qrcodeKey key = "qrcode"
const jdToken key = "jd_token"

// WithUserProfile 往已有的context中注入管理员信息
func WithUserProfile(ctx context.Context, admin *pkg.User) context.Context {
	return context.WithValue(ctx, userKey, admin)
}

// GetUserProfile 从context中提取管理员
func GetUserProfile(ctx context.Context) (*pkg.User, bool) {
	admin, ok := ctx.Value(userKey).(*pkg.User)
	return admin, ok
}

// WithQrcode 往已有的context中注入qrcodeParam
func WithQrcode(ctx context.Context, qrcode *pkg.QrcodeParam) context.Context {
	return context.WithValue(ctx, qrcodeKey, qrcode)
}

// GetQrcode 从context中提取qrcodeParam
func GetQrcode(ctx context.Context) (*pkg.QrcodeParam, bool) {
	qrcode, ok := ctx.Value(qrcodeKey).(*pkg.QrcodeParam)
	return qrcode, ok
}

// WithJdToken 往已有的context中注入qrcodeParam
func WithJdToken(ctx context.Context, Token *pkg.QrcodeJdToken) context.Context {
	return context.WithValue(ctx, jdToken, Token)
}

// GetJdToken 从context中提取qrcodeParam
func GetJdToken(ctx context.Context) (*pkg.QrcodeJdToken, bool) {
	token, ok := ctx.Value(jdToken).(*pkg.QrcodeJdToken)
	return token, ok
}
