package main

import (
	"strings"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	p := &Page{HeaderMessage{Visible: "hidden"}, false, nil}	
	if isSession(r) {
		http.Redirect(w, r, "/servers", 301)
	}
	r.ParseForm()
	if !checkPost(r.PostForm, "username", "password") {
		loginPage(w, r, p)
		return
	}
	user := strings.TrimSpace(r.PostFormValue("newusername"))
	pass := strings.TrimSpace(r.PostFormValue("password"))
	if checkAccount(user, pass) {
		addSession(w)
		http.Redirect(w, r, "/servers", 301)
	}
}

func signup(w http.ResponseWriter, r *http.Request) {
	p := &Page{HeaderMessage{Visible: "hidden"}, false, nil}	
	if isSession(r) {
		http.Redirect(w, r, "/servers", 301)
	}
	r.ParseForm()
	if !checkPost(r.PostForm, "newusername", "password", "confpassword") {
		signupPage(w, r, p)
		return
	}
	user := strings.TrimSpace(r.PostFormValue("newusername"))
	pass := strings.TrimSpace(r.PostFormValue("password"))
	if err := addAdmin(user, pass); err == nil {
		addSession(w)
		http.Redirect(w, r, "/servers", 301)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	endSession(w, r)
}

func servers(w http.ResponseWriter, r *http.Request) {
	p := &Page{HeaderMessage{Visible: "hidden"}, true, nil}
			// 	p.Mess.Type = "danger"
			// p.Mess.Message = "Please fill out all fields"
			// p.Mess.Visible = ""

	if !isSession(r) {
		http.Redirect(w, r, "/login", 301)
	}
	serversPage(w, r, p)
}
