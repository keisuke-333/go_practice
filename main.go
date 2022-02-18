package main

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/keisuke-333/go_practice/entity"
)

func main() {
	// ————— 環境変数の読み込み —————
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// —————————————————————————

	// ————— MySQL接続 —————
	db, err := gorm.Open("mysql", os.Getenv("CONNECT"))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()
	db.LogMode(true)
	if err := db.DB().Ping(); err != nil {
		panic(err)
	}
	// —————————————————————————

	// ————— テーブル作成 —————
	db.AutoMigrate(
		&entity.User{},
		&entity.Organization{},
	)
	// —————————————————————————
}
