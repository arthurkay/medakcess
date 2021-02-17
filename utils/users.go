package utils

import (
	"medakcess/models"
	"net/http"

	"gorm.io/gorm"
)

func GetAllUsers(db *gorm.DB, r *http.Request) ([]models.User, error) {
	var (
		users []models.User
	)

	results := db.Find(&users)

	if results.RowsAffected < 1 {
		return users, results.Error
	} else {
		return users, nil
	}
}

func GetUserTypes() []models.UserType {
	db, _ := models.DBConfig()
	var userTypes []models.UserType
	db.Find(&userTypes)
	return userTypes
}
