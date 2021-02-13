package main

import (
	"medakcess/controllers"
	"medakcess/middleware"
	"medakcess/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.StrictSlash(true)

	// serving static files
	staticFileDirectory := http.Dir(utils.AppFilePath("assets/"))

	staticFileHandler := http.StripPrefix("/assets/",
		http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	// serve a file
	r.HandleFunc("/favicon.ico", controllers.ServeFileHandler)
	r.HandleFunc("/robots.txt", controllers.ServeFileHandler)
	r.HandleFunc("/ads.txt", controllers.ServeFileHandler)

	// Auth Routes
	auth := r.NewRoute().Subrouter()
	// auth.Use(middleware.AuthPagesAccess)
	auth.HandleFunc("/login", controllers.LoginHandler).Methods("POST")
	
	// Admin auth routes
	admin := r.PathPrefix("/admin").Subrouter()
	admin.Use(middleware.AuthAdmin)
	admin.HandleFunc("/", controllers.AdminHome).Methods("GET")
	admin.HandleFunc("/users", controllers.Users).Methods("GET")
	admin.HandleFunc("/user", controllers.CreateUser).Methods("GET", "POST")
	admin.HandleFunc("/user/:id", controllers.UpdateUser).Methods("GET", "PUT","DELETE")

	// Dashboard auth routes
	dash := r.PathPrefix("/dashboard").Subrouter()
	// dash.Use(middleware.AuthPagesAccess)
	dash.Use(middleware.Authorised)
	dash.HandleFunc("/", controllers.DashHomeHandler).Methods("GET")

	// controllers Routes
	r.HandleFunc("/", controllers.HomeHandler).Methods("GET")

	return r
}
