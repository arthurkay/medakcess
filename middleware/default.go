package middleware

import (
	"log"
	"net/http"
	"medakcess/utils"
)

// Only show login/signup routes if not logged in
func AuthPagesAccess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Do not display login/register pages if user is already logged in")
		uid, err := utils.GetUserID(r)
		if err == nil {
			// User is logged in, do not show pages
			log.Printf("Possibly user has a cookie. %s", uid)
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
		} else {
			// User is not logged in, show pages
			log.Printf("Ohh snap! no cookies for ya!, %s", err.Error())
			next.ServeHTTP(w, r)
		}

	})
}
