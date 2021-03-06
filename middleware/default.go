package middleware

import (
	"log"
	"medakcess/utils"
	"net/http"
)

// AuthPagesAccess Only show login routes if not logged in
func AuthPagesAccess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uid, err := utils.GetUserID(r)
		log.Printf("UserID: %d", uid)
		if err == nil {
			// User is logged in, do not show pages
			log.Printf("User has a cookie. %d", uid)
			next.ServeHTTP(w, r)
			return
		} else {
			// User is not logged in, show pages
			log.Printf("Ohh snap! no cookies for ya!, %s", err.Error())
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

	})
}

// Authorised checks whether or not the user is authenticated
func Authorised(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := utils.GetUserID(r)
		if err != nil {
			// User has no session, take them to home page
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

// AuthAdmin authorise admin access
func AuthAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userType, err := utils.GetUserType(r)
		if err == nil {
			if userType == 1 {
				next.ServeHTTP(w, r)
			} else {
				http.Redirect(w, r, "/dashboard/", http.StatusFound)
			}
		} else {
			http.Redirect(w, r, "/", http.StatusFound)
		}
	})
}
