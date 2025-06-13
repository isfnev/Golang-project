package main

import (
	"net/http"
	"fmt"
	"time"
)

type Login struct{
	HashedPassword string
	SessionToken string
	CSRFToken string
}

var users = map[string]Login{}

func main(){
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/protected", protected)

	http.ListenAndServe(":8080", nil)
}

func register(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	if len(username)<8 || len(password)<8 {
		http.Error(w, "Invalid username/password", http.StatusNotAcceptable)
		return
	}

	if _, ok := users[username]; ok {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	hashPassword, _ := hashPassword(password)
	users[username] = Login{
		HashPassword : hashPassword,
	}
}

func login(w http.ResponseWriter, r *http.Request){
	
}

func logout(w http.ResponseWriter, r *http.Request){
	
}

func protected(w http.ResponseWriter, r *http.Request){
	
}