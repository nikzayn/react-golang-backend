package routing

import (
	"fmt"
	"net/http"

	helpers "../helpers"
	user "../user"
	"github.com/gorilla/securecookie"
)

//Random cookie generator
var cookieHandler = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

//Route Handlers

//GET Request for login page
func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := helpers.LoadFile("templates/login.html")
	fmt.Fprintf(w, body)
}

//POST Request for login page
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	pass := r.FormValue("password")
	redirectTarget := "/"

	//if helpers are not empty with(name) and (password)
	if !helpers.IsEmpty(name) && !helpers.IsEmpty(password) {
		//Database check for user data
		_userIsValid := user.UserIsValid(name, pass)

		if _userIsValid {
			SetCookie(name, w)
			redirectTarget = "/index"
		} else {
			redirectTarget = "/register"
		}
	}
	http.Redirect(w, r, redirectTarget, 302)	
}


//GET Request for register page
func RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := helpers.LoadFile("templates/register.html")
	fmt.Fprintf(w, body)
}

//POST Request for register page
func RegisterHandler(w http.ResponseWriter, r *http.Request){
	//Parse the form
	r.ParseForm()

	uName := r.FormValue("username")
	email := r.FormValue("email")
	pwd := r.FormValue("password")
	confirmPwd := r.FormValue("confirmPassword")

	//If the username, email, password and confirm password are true
	if(uName && email && pwd && confirmPwd){
		fmt.Fprintln(w, "Username for Register : ", uName)
		fmt.Fprintln(w, "Email for Register : ", email)
		fmt.Fprintln(w, "Password for Register : ", pwd)
		fmt.Fprintln(w, "confirmPwd for Register : ", confirmPwd)
	}
	else{
		fmt.Fprintln(w, "This fields can not be blank!")
	}
}

//GET Request for Index Page
func IndexPageHandler(w http.ResponseWriter, r *http.Request){
	username := GetUserName
}

//POST Request for Logout Page
func LogoutHandler(w http.ResponseWriter, r *http.Request){
	ClearCookie(w)
	http.Redirect(w, r, "/", 302)
}