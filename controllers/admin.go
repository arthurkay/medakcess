package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"medakcess/models"
	"medakcess/utils"
	"net/http"
	"strconv"
)

func AdminHome(w http.ResponseWriter, r *http.Request) {
	data := utils.TemplateRenderer{
		View:      utils.AppFilePath("templates/admin/home.html"),
		Hview:     utils.AppFilePath("templates/layouts/header.tpl"),
		Fview:     utils.AppFilePath("templates/layouts/footer.tpl"),
		PageTitle: "Admin Home | MedAkcess",
		UserName:  GetUserName(r),
	}

	tmpl := template.Must(template.ParseFiles(data.View, data.Hview, data.Fview))
	tmpl.Execute(w, data)
}

func Users(w http.ResponseWriter, r *http.Request) {
	var data utils.TemplateRenderer
	users, error := utils.GetAllUsers(db, r)
	dataSlice := make(map[string]interface{})

	if error != nil {
		fmt.Print("Unable to get users")
	} else {
		dataSlice["users"] = users
		dataSlice["usertypes"] = utils.GetUserTypes()

		data = utils.TemplateRenderer{
			View:      utils.AppFilePath("templates/admin/users.html"),
			Hview:     utils.AppFilePath("templates/layouts/header.tpl"),
			Fview:     utils.AppFilePath("templates/layouts/footer.tpl"),
			PageTitle: "Platform  Users | MedAkcess",
			UserName:  GetUserName(r),
			Data:      dataSlice,
		}
	}
	tmpl := template.Must(template.ParseFiles(data.View, data.Hview, data.Fview))
	tmpl.Execute(w, data)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	data := utils.TemplateRenderer{
		View:      utils.AppFilePath("templates/admin/create-user.html"),
		Hview:     utils.AppFilePath("templates/layouts/header.tpl"),
		Fview:     utils.AppFilePath("templates/layouts/footer.tpl"),
		PageTitle: "Admin Home | MedAkcess",
		UserName:  GetUserName(r),
	}

	if r.Method == "POST" {
		db, err := models.DBConfig()
		if err != nil {
			panic("System unable to get database instance")
		} else {
			password, _ := utils.BcryptHash("Test@321")
			usertype, _ := strconv.Atoi(r.PostForm.Get("usertype"))
			db.Create(&models.User{
				FirstName:  r.PostForm.Get("fname"),
				MiddleName: r.PostForm.Get("mname"),
				LastName:   r.PostForm.Get("lname"),
				Email:      r.PostForm.Get("email"),
				Password:   string(password),
				UserTypeID: uint(usertype),
			})
		}
	} else {
		tmpl := template.Must(template.ParseFiles(data.View, data.Hview, data.Fview))
		tmpl.Execute(w, data)
	}

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	data := utils.TemplateRenderer{
		View:      utils.AppFilePath("templates/admin/update-user.html"),
		Hview:     utils.AppFilePath("templates/layouts/header.tpl"),
		Fview:     utils.AppFilePath("templates/layouts/footer.tpl"),
		PageTitle: "Admin Home | MedAkcess",
		UserName:  GetUserName(r),
	}

	tmpl := template.Must(template.ParseFiles(data.View, data.Hview, data.Fview))
	tmpl.Execute(w, data)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := utils.GetAllUsers(db, r)
	if err != nil {
		fmt.Print("Unable to get users")
	} else {
		data, err := json.Marshal(users)
		if err == nil {
			n := len(data)
			s := string(data[:n])
			log.Printf("%s", string(data))
			fmt.Fprintf(w, s)
		} else {
			fmt.Println("")
		}
	}
}
