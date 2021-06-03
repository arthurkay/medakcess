package models

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConfig() (*gorm.DB, error) {
	godotenv.Load()
	var (
		host     string = os.Getenv("DB_HOST")
		user     string = os.Getenv("DB_USER")
		password string = os.Getenv("DB_PASSWORD")
		dbname   string = os.Getenv("DB_NAME")
	)
	db_config := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s 
                        sslmode=disable`,
		host,
		user,
		password,
		dbname)

	db, err := gorm.Open(postgres.Open(db_config), &gorm.Config{})
	return db, err
}

func DBMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
		&UserType{},
		&Patient{},
		&NextOfKeen{},
		&Vitals{},
		&Triage{},
		&Booking{},
		&Specimen{},
		&Blood{},
		&FBC{},
		&Urine{},
		&Urinalysis{},
		&UrineMCS{},
		&Stool{},
		&StoolMCS{},
		&Other{},
	)
}

func DBSeed(db *gorm.DB) {
	db.Create(&UserType{
		Name: "Admin",
	})

	db.Create(&UserType{
		Name: "Registry",
	})

	db.Create(&UserType{
		Name: "Nurse",
	})

	db.Create(&UserType{
		Name: "Clinician",
	})

	passwordHash, passwordErr := bcrypt.GenerateFromPassword([]byte("P@55w0rd"), bcrypt.DefaultCost)

	if passwordErr == nil {
		db.Create(&User{
			FirstName:  "Arthur",
			LastName:   "Kalikiti",
			Email:      "arthur@kalikiti.net",
			Password:   string(passwordHash),
			UserTypeID: 1,
		})
	} else {
		panic("Unable to create password hash")
	}
}
