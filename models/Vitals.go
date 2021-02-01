package models

import (
	"gorm.io/gorm"
)

type Vitals struct {
	gorm.Model
	UserID    uint
	User      User
	PatientID uint
	Patient   Patient
	// Temperature in degrees celcius
	Temperature float64
	// Blood pressure mm/Hg
	BP int
	// Pulse bit/minute
	Pulse int
	// Weight in Kg's
	Weight float64
	// Bload Sugar, 0 fasting, 1 random
	BS int `gorm:"type:int;enum(0,1)"`
	// Blood sugar values grams/ltr
	BSValue float64
}

type Triage struct {
	gorm.Model
	VitalsID uint
	Vitals   Vitals
	// 0 Mild, 1 Moderate, 2 Severe
	Scale   int `gorm:"type:int;enum(0,1,2)"`
	Comment string
}
