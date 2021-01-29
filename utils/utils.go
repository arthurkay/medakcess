package utils

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/securecookie"
)

// Session handling
var cookie_handler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func AppFilePath(fpath string) string {
	appPath, err := os.Getwd()
	if err != nil {
		log.Printf("Unable to get current working application path")
	}
	return filepath.Join(appPath, fpath)
}

func GetUserID(request *http.Request) (string, error) {
	userid := ""
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		err = cookie_handler.Decode("session", cookie.Value, &cookieValue)
		userid = cookieValue["uid"]
		if err != nil {
			return userid, fmt.Errorf("No user session stored in cookie!")
		}
	} else {
		return userid, fmt.Errorf("No user session stored in cookie!")
	}
	// fmt.Println(userid)
	return userid, nil
}

// Template handler & rendering
type TemplateRenderer struct {
	View      string      // Main body of page
	Hview     string      // Header of page
	Fview     string      // Footer of the rendered page
	PageTitle string      // Web page title
	UserName  string      // Logged In User
	Data      interface{} // Misc data
}
