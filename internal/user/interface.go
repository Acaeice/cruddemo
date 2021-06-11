package user

import (
	"context"

	"code.meikeland.com/wanghejun/cruddemo/pkg"
)

type userRepoI interface {
	Create(param pkg.User) (*pkg.User, error)
	GetCode(ctx context.Context, code string) (*pkg.Code2Session, error)
	GetUser(sessionKey, encryptedData, iv string) (*pkg.MAppUser, error)
	GetByID(id uint) (*pkg.User, error)
	GetByOpenid(openid string) (*pkg.User, error)
}
