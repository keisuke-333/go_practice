package main

import (
	"sandbox/infra"
	"sandbox/models"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()
	if err := db.AutoMigrate(&models.Item{}); err != nil {
		panic("Faild to migrate database")
	}
}
