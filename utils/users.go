package utils

import (
	"fmt"
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

func CreateUser(
	fname string, mname string, lname string,
	email string, password string, usertype uint) (string, error) {
	pass, err := BcryptHash(password)
	if err != nil {
		db, _ := models.DBConfig()
		db.Create(&models.User{
			FirstName:  fname,
			MiddleName: mname,
			LastName:   lname,
			Email:      email,
			Password:   string(pass),
			UserTypeID: usertype,
		})
		return "ok", nil
	} else {
		return "", fmt.Errorf("Unable to create user")
	}
}
