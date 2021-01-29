package main

import (
	"medakcess/models"

	"gorm.io/gorm"
)

func DBMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.UserType{},
	)
}
