package main

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func getUser(w http.ResponseWriter, r *http.Request) user {
	//get cookie
	c, err := r.Cookie("session")
	if err != nil {
		id, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: id.String(),
		}
	}
	http.SetCookie(w, c)

	//If user already exists
	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}
	return u
}

//If users is already logged in
func alreadyLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}
