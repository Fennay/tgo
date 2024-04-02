package store

import (
	"github.com/fennay/tgo/internal/utils/db"
	"gorm.io/gorm"
)

func Master() *gorm.DB {
	return db.Init()
}
