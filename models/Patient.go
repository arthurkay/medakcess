package models

import (
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	FirstName    string `gorm:"type: varchar(100); not null"`
	MiddleName   string `gorm:"type: varchar(100)"`
	LastName     string `gorm:"type: varchar(100); not null"`
	Age          int
	Contact      int
	NextOfKeenID uint
	NextOfKeen   NextOfKeen
}

type NextOfKeen struct {
	gorm.Model
	FirstName  string `gorm:"type: varchar(100); not null"`
	MiddleName string `gorm:"type: varchar(100)"`
	LastName   string `gorm:"type: varchar(100); not null"`
	Age        int
	Contact    int
}
