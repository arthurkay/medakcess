package main

import (
	"log"
	"medakcess/models"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	db, dberr := models.DBConfig()

	if dberr != nil {
		log.Printf("Unable to connect to database: %s", dberr)
	}

	// Check if user has passed extra parameters
	if len(os.Args) > 1 {
		args := os.Args[1:]
		var i int
		for i = 0; i < len(args); i++ {
			// If user passed migrate command, migrate database schema
			if args[i] == "migrate" {
				// Migrate Database
				DBMigrate(db)
			}

			if args[i] == "seed" {
				DBSeed(db)
			}
		}
		// Exit process after migrating/seeding data
		os.Exit(0)
	}

	// Load .env file
	godotenv.Load()
	// Load app routes
	routes := Router()
	var port string
	port = os.Getenv("PORT")
	log.Printf("Serving on port %s", port)
	var err error = http.ListenAndServe(":"+port, routes)
	if err != nil {
		log.Fatal("Oops, something went wrong " + err.Error())
	}

}
