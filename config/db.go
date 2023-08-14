package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var dbDefault *gorm.DB

// sử dụng singleton pattern để tạo một connection duy nhất đến database
// khi ứng dụng lớn hơn thì không nên sử dụng singleton pattern
// thay vào đó nên sử dụng connection pool
func (a *App) GetDB() *gorm.DB {
	if dbDefault == nil {
		return a.initDB()
	}
	return dbDefault
}

func (a *App) initDB() *gorm.DB {
	// Tạo chuỗi kết nối đến PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	// Kết nối đến cơ sở dữ liệu PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	return db
}
