package main

import (
	"strings"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	if isSession(r) {
		http.Redirect(w, r, "/home", 301)
	}
	loginPage(w, r)
	if !checkPost(r.PostForm, "username", "password") {
		return
	}
	user := strings.TrimSpace(r.PostFormValue("newusername"))
	pass := strings.TrimSpace(r.PostFormValue("password"))
	if checkAccount(user, pass) {
		addSession(w)
		http.Redirect(w, r, "/home", 301)
	}
}

func signup(w http.ResponseWriter, r *http.Request) {
	if isSession(r) {
		http.Redirect(w, r, "/home", 301)
	}
	signupPage(w, r)
	if !checkPost(r.PostForm, "newusername", "password", "confpassword") {
		return
	}
	user := strings.TrimSpace(r.PostFormValue("newusername"))
	pass := strings.TrimSpace(r.PostFormValue("password"))
	if err := addAdmin(user, pass); err == nil {
		addSession(w)
		http.Redirect(w, r, "/home", 301)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	endSession(w, r)
}

func servers(w http.ResponseWriter, r *http.Request) {

}
