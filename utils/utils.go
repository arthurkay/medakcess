package utils

import (
	"fmt"
	"log"
	"medakcess/models"
	"net/http"
	"os"
	"path/filepath"

	"gorm.io/gorm"

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

func SetCookieHandler(w http.ResponseWriter, r *http.Request, user models.User) {
	value := map[string]uint{
		"session": user.ID,
	}
	if encoded, err := cookie_handler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}

		http.SetCookie(w, cookie)
	}
}

func GetUserID(r *http.Request) (uint, error) {
	var userid uint
	if cookie, err := r.Cookie("session"); err == nil {
		cookieValue := make(map[string]uint)
		err = cookie_handler.Decode("session", cookie.Value, &cookieValue)
		userid = cookieValue["session"]
		if err != nil {
			return userid, fmt.Errorf("No user session stored in cookie!")
		}
	} else {
		return userid, fmt.Errorf("No user session stored in cookie!")
	}
	fmt.Println(userid)
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

func FindUserByEmail(db *gorm.DB, email string) (*models.User, error) {
	var user models.User
	result := db.Where("Email = ?", email).Find(&user)
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("Unable to find user with %s email", email)
	} else {
		return &user, nil
	}
}

func GetUserType(r *http.Request) (uint, error) {
	uid, err := GetUserID(r)
	if err != nil {
		// uid = 0
		return uid, fmt.Errorf("User not found")
	} else {
		db, _ := models.DBConfig()
		var user models.User
		db.Find(&user, uid)
		log.Printf("user is of type: %s", user.UserTypeID)
		return user.UserTypeID, nil
	}
}
