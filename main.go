package main

import (
	"log"
	// "net/http"
	"os"

	// "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/keisuke-333/go_practice/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// ————— 環境変数の読み込み —————
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// —————————————————————————

	// ————— MySQL接続 —————
	dns := os.Getenv("CONNECT")
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
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
	db.AutoMigrate(
		&entity.User{},
	)
	// —————————————————————————
}
