package models

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
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
