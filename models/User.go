package models

import (
	"gorm.io/gorm"
)

type UserType struct {
	gorm.Model
	Name string `gorm:"type: varchar(50);not null;unique_index"`
}

type User struct {
	gorm.Model
	FirstName  string `gorm:"type: varchar(150);not null"`
	MiddleName string
	LastName   string `gorm:"type: varchar(150);not null"`
	Email      string `gorm:"type: varchar(100);unique_index;not null"`
	Password   string
	UserTypeID uint `gorm:"not null"`
	UserType   UserType
}
