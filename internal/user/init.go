package user

import "code.meikeland.com/wanghejun/cruddemo/internal/sql"

var (
	UserRepo userRepoI
)

func Init() {
	UserRepo = userSQLRepo{
		db: sql.Db,
	}
}
