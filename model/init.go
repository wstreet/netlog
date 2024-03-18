package model

import (
	"netlog/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	c := config.GetConfig()

	db, err := gorm.Open(sqlite.Open(c.Sqlite.Filename), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
	// Migrate the schema
	db.AutoMigrate(&Item{})
}
