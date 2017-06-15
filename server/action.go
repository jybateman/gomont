package main

import (
	"strings"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	loginPage(w, r)
}

func signup(w http.ResponseWriter, r *http.Request) {
	signupPage(w, r)
	if !checkPost(r.PostForm, "newusername", "password", "confpassword") {
		return
	}
	user := strings.TrimSpace(r.PostFormValue("newusername"))
	pass := strings.TrimSpace(r.PostFormValue("password"))
	addAdmin(user, pass)
}

func logout(w http.ResponseWriter, r *http.Request) {
	loginPage(w, r)
	if !checkPost(r.PostForm, "username", "password") {
		return
	}
	user := strings.TrimSpace(r.PostFormValue("newusername"))
	pass := strings.TrimSpace(r.PostFormValue("password"))
	checkAccount(user, pass)
}

func servers(w http.ResponseWriter, r *http.Request) {

}
