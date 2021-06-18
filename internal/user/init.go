package user

import "github.com/wechatapi/cruddemo/internal/sql"

var (
	UserRepo userRepoI
)

func Init() {
	UserRepo = userSQLRepo{
		db: sql.Db,
	}
}
