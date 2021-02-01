package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	UserID    uint
	User      User
	PatientID uint
	Patient   Patient
	// 0 Clinician, 1 Lab, 3 Radiology
	BookingDepartment int       `gorm:"type:int;enum(0,1,2)"`
	TimeStamp         time.Time `gorm:"type:time;not null"`
}
