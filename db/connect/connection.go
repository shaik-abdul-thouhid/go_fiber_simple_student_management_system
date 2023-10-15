package connect

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn = "user:password@tcp(127.0.0.1:3306)/employee_directory?charset=utf8mb4&parseTime=True&loc=Local"

var db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

func GetDB() (*gorm.DB, error) {
	return db, err
}
