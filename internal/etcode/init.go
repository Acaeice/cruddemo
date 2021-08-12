package etcode

import "github.com/wechatapi/cruddemo/internal/sql"

var (
	QrcodeRepo qrCodeRepoI
)

func Init() {
	QrcodeRepo = qrCodeSQLRepo{
		db: sql.Db,
	}
}
