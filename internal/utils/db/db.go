package db

import (
	"fmt"
	"github.com/fennay/tgo/internal/utils/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() *gorm.DB {
	if db != nil {
		return db
	}

	dbConfig := config.GetConfig().DB

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.URL, dbConfig.Port, dbConfig.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
