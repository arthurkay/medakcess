package main

import (
	"medakcess/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func DBMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.UserType{},
		&models.Patient{},
		&models.NextOfKeen{},
		&models.Vitals{},
		&models.Triage{},
		&models.Booking{},
		&models.Specimen{},
		&models.Blood{},
		&models.FBC{},
		&models.Urine{},
		&models.Urinalysis{},
		&models.UrineMCS{},
		&models.Stool{},
		&models.StoolMCS{},
		&models.Other{},
	)
}

func DBSeed(db *gorm.DB) {
	db.Create(&models.UserType{
		Name: "Registry",
	})

	db.Create(&models.UserType{
		Name: "Nurse",
	})

	db.Create(&models.UserType{
		Name: "Clinician",
	})

	passwordHash, passwordErr := bcrypt.GenerateFromPassword([]byte("P@55w0rd"), bcrypt.DefaultCost)

	if passwordErr == nil {
		db.Create(&models.User{
			FirstName:  "Arthur Kalikiti",
			LastName:   "Kalikiti",
			Email:      "arthur@kalikiti.net",
			Password:   string(passwordHash),
			UserTypeID: 3,
		})
	}
}
