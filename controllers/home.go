package controllers

import (
	"html/template"
	"medakcess/utils"
	"net/http"
	"path"
)

func ServeFileHandler(res http.ResponseWriter, req *http.Request) {
	fname := path.Base(req.URL.Path)
	filePath := "assets/" + fname
	http.ServeFile(res, req, utils.AppFilePath(filePath))
}

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	data := utils.TemplateRenderer{
		View:      utils.AppFilePath("templates/home.html"),
		Hview:     utils.AppFilePath("templates/layouts/header.tpl"),
		Fview:     utils.AppFilePath("templates/layouts/footer.tpl"),
		PageTitle: "MedAkcess | Welcome",
	}

	tmpl := template.Must(template.ParseFiles(data.View, data.Hview, data.Fview))
	tmpl.Execute(res, data)
}
