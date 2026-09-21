package config

import (
	"github.com/AKSHAYK0UL/graph-ql/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	DB.AutoMigrate(&models.User{})

	return nil
}
