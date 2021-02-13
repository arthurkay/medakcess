package controllers

import (
	"fmt"
	"html/template"
	"medakcess/utils"
	"net/http"
)

func AdminHome(w http.ResponseWriter, r *http.Request) {
	data := utils.TemplateRenderer{
		View:      utils.AppFilePath("templates/admin/home.html"),
		Hview:     utils.AppFilePath("templates/layouts/header.tpl"),
		Fview:     utils.AppFilePath("templates/layouts/footer.tpl"),
		PageTitle: "Admin Home | MedAkcess",
	}

	tmpl := template.Must(template.ParseFiles(data.View, data.Hview, data.Fview))
	tmpl.Execute(w, data)
}

func Users(w http.ResponseWriter, r *http.Request) {
	data := utils.TemplateRenderer{
		View:      utils.AppFilePath("templates/admin/users.html"),
		Hview:     utils.AppFilePath("templates/layouts/header.tpl"),
		Fview:     utils.AppFilePath("templates/layouts/footer.tpl"),
		PageTitle: "Platform  Users | MedAkcess",
	}

	users, error := utils.GetAllUsers(db, r)
	if error != nil {
		fmt.Print("Unable to get users")
	} else {
		for _, v := range users {
			fmt.Println(v)
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
	}

	tmpl := template.Must(template.ParseFiles(data.View, data.Hview, data.Fview))
	tmpl.Execute(w, data)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	data := utils.TemplateRenderer{
		View:      utils.AppFilePath("templates/admin/update-user.html"),
		Hview:     utils.AppFilePath("templates/layouts/header.tpl"),
		Fview:     utils.AppFilePath("templates/layouts/footer.tpl"),
		PageTitle: "Admin Home | MedAkcess",
	}

	tmpl := template.Must(template.ParseFiles(data.View, data.Hview, data.Fview))
	tmpl.Execute(w, data)
}
