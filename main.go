package main

import (
	"log"

	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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

	r := gin.Default()

	r.GET("/users", func(c *gin.Context) {
		users := []entity.User{}
		db.Find(&users)
		c.JSON(http.StatusOK, users)
	})

	r.GET("/user/:id", func(c *gin.Context) {
		user := entity.User{}
		if err := db.Where("ID = ?", c.Param("id")).First(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": user})
	})

	r.POST("/user", func(c *gin.Context) {
		user := entity.User{}
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": user})
	})

	r.PUT("/user/:id", func(c *gin.Context) {
		user := entity.User{}
		if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
		input := entity.User{}
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Where("ID = ?", c.Param("id")).First(&user).Updates(&input)
		c.JSON(http.StatusOK, gin.H{"data": user})
	})

	r.DELETE("/user/:id", func(c *gin.Context) {
		user := entity.User{}
		if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
		db.Delete(&user)
		c.JSON(http.StatusOK, gin.H{"data": true})
	})

	r.Run(":8080")
}
