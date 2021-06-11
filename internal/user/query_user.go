package user

import "gorm.io/gorm"

// subscribeCountQuery subscribeCount查询条件
type User struct {
	ID     uint
	OpenID string
	Limit  uint
	Offset uint
}

// where 条件
func (c User) where() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if c.ID > 0 {
			db = db.Where("id = ?", c.ID)
		}

		return db
	}
}

// order 处理排序规则
func (c User) order() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Order("created_at DESC")
		return db
	}
}

// preload 预加载表
func (c User) preload() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// if c.PreloadProperty {
		// 	db = db.Preload("PropertyArr")
		// }
		return db
	}
}
