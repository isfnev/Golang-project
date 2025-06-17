package main

import (
	"errors"
	"fmt"
	"net/http"
)

var ErrAuth = errors.New("Unauthorized")

func Authorize(w http.ResponseWriter, r *http.Request) error {
	username := r.FormValue("username")
	user, ok := users[username]
	fmt.Fprintln(w, "username : ", username)
	fmt.Fprintln(w, "user session token : ", user.SessionToken, " and csrf token : ", user.CSRFToken)

	if !ok {
		panic(ErrAuth)
	}

	st, err := r.Cookie("session_token")
	// fmt.Fprintln(w, "session token : ", st, " and ", err)
	if err != nil || st.Value != user.SessionToken {
		panic(err)
	}

	csrf := r.Header.Get("X-CSRF-Token")
	if csrf != user.CSRFToken {
		panic(ErrAuth)
	}
	return nil
}
