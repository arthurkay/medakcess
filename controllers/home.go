package controllers

import (
	"html/template"
	"log"
	"medakcess/utils"
	"net/http"
	"path"

	"golang.org/x/crypto/bcrypt"
)

func ServeFileHandler(w http.ResponseWriter, r *http.Request) {
	fname := path.Base(r.URL.Path)
	filePath := "assets/" + fname
	http.ServeFile(w, r, utils.AppFilePath(filePath))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := utils.TemplateRenderer{
		View:      utils.AppFilePath("templates/home.html"),
		Hview:     utils.AppFilePath("templates/layouts/header.tpl"),
		Fview:     utils.AppFilePath("templates/layouts/footer.tpl"),
		PageTitle: "MedAkcess | Welcome",
	}

	tmpl := template.Must(template.ParseFiles(data.View, data.Hview, data.Fview))
	tmpl.Execute(w, data)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var (
			email    string = r.FormValue("email")
			password string = r.FormValue("password")
			// userModel *models.User
		)

		if dberr == nil {
			userModel, dataErr := utils.FindUserByEmail(db, email)
			if dataErr == nil {
				passErr := bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(password))
				if passErr == nil {
					log.Printf("Successful log in by %s %s", userModel.FirstName, userModel.LastName)
					utils.SetCookieHandler(w, r, *userModel)
					http.Redirect(w, r, "/dashboard", http.StatusFound)
				} else {
					log.Print("Incorrect username/password combination")
					http.Redirect(w, r, "/", http.StatusFound)
				}
			} else {
				log.Print("Incorrect username/password combination")
				http.Redirect(w, r, "/", http.StatusFound)
			}
		} else {
			log.Printf("Unable to make database connection: %s", dberr)
		}
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func DashHomeHandler(w http.ResponseWriter, r *http.Request) {
	data := utils.TemplateRenderer{
		View:      utils.AppFilePath("templates/dash/home.html"),
		Hview:     utils.AppFilePath("templates/layouts/header.tpl"),
		Fview:     utils.AppFilePath("templates/layouts/footer.tpl"),
		PageTitle: "Dashboard Home | MedAkcess",
	}

	tmpl := template.Must(template.ParseFiles(data.View, data.Hview, data.Fview))
	tmpl.Execute(w, data)
}
