package controllers

import (
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
