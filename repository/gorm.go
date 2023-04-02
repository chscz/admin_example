package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db = InitDatabase()

func InitDatabase() *gorm.DB {
	dsn := "root:1111@tcp(127.0.0.1:3306)/admin?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("database connect error")
	}
	return db
}
