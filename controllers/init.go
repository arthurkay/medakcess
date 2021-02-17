package controllers

import (
	"medakcess/models"
	"medakcess/utils"

	"net/http"

	"gorm.io/gorm"
)

var (
	db    *gorm.DB
	dberr error
)

func init() {
	db, dberr = models.DBConfig()
	if dberr != nil {
		panic("Hoops, unable to get db access")
	}
}

func GetUserName(r *http.Request) string {
	userid, error := utils.GetUserID(r)
	if error == nil {
		user, err := utils.FindUserById(db, userid)
		if err == nil {
			return user.FirstName + " " + user.LastName
		}
	}
	return ""
}
