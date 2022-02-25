package model

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect() {
	// ————— 環境変数の読み込み —————
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	// —————————————————————————

	// ————— MySQL接続 —————
	dns := os.Getenv("DB_DNS")
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	defer func() {
		sqlDB, err := db.DB()
		if err == nil {
			sqlDB.Close()
		}
	}()
	// —————————————————————————

	// ————— SQLログ出力 —————
	db.Logger = db.Logger.LogMode(4)
	// —————————————————————————

	// ————— テーブル作成 —————
	// db.AutoMigrate(
	// 	Task{},
	// )
	// —————————————————————————

	DB = db
}
