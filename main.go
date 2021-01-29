package main

import (
	"log"
	"medakcess/models"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	godotenv.Load()
	// Load app routes
	routes := Router()
	var port string
	port = os.Getenv("PORT")
	log.Printf("Serving on port %s", port)
	err := http.ListenAndServe(":"+port, routes)
	if err != nil {
		log.Fatal("Oops, something went wrong " + err.Error())
	}

	db, err := models.DBConfig()

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
	}

}
