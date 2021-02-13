package controllers

import (
	"medakcess/models"

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
