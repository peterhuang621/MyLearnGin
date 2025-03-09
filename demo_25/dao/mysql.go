package dao

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitMySQL() error {
	var err error
	dsn := "root:@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("failed to get database connection: %v", err)
	}

	return sqlDB.Ping()
}
