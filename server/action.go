package main

import (
	"fmt"
	"strings"
	"strconv"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	p := &Page{HeaderMessage{Visible: "hidden"}, false, nil}
	if isSession(r) {
		http.Redirect(w, r, "/servers", 302)
	}
	r.ParseForm()
	if !checkPost(r.PostForm, "username", "password") {
		loginPage(w, r, p)
		return
	}
	user := strings.TrimSpace(r.PostFormValue("username"))
	pass := strings.TrimSpace(r.PostFormValue("password"))
	if checkAccount(user, pass) {
		addSession(w)
		http.Redirect(w, r, "/servers", 302)
	}
	p.Mess.Type = "Warning"
	p.Mess.Message = "Wrong Username/Password"
	p.Mess.Visible = ""
	loginPage(w, r, p)
}

func signup(w http.ResponseWriter, r *http.Request) {
	p := &Page{HeaderMessage{Visible: "hidden"}, false, nil}
	if isSession(r) {
		http.Redirect(w, r, "/servers", 302)
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
		http.Redirect(w, r, "/servers", 302)
	}
	p.Mess.Type = "Danger"
	p.Mess.Message = "Couldn't create account"
	p.Mess.Visible = ""
	signupPage(w, r, p)
}

func logout(w http.ResponseWriter, r *http.Request) {
	endSession(w, r)
	http.Redirect(w, r, "/login", 302)
}

func servers(w http.ResponseWriter, r *http.Request) {
	var s Servers
	var err error

	p := &Page{HeaderMessage{Visible: "hidden"}, true, nil}
	//	p.Mess.Type = "danger"
	// p.Mess.Message = "Please fill out all fields"
	// p.Mess.Visible = ""

	if !isSession(r) {
		http.Redirect(w, r, "/login", 302)
	}
	s.Srvs, err = getServer()
	if err != nil {
		fmt.Println(err)
		p.Mess.Type = "Danger"
		p.Mess.Message = "Couldn't get Servers"
		p.Mess.Visible = ""
	}
	p.Info = s
	serversPage(w, r, p)
}

func server(w http.ResponseWriter, r *http.Request) {
	var s Servers
	var err error

	r.ParseForm()
	if checkPost(r.PostForm, "_method") && r.PostForm["_method"][0] == "delete" {
		id := r.URL.Path[len("/server/"):]
		delServer(id)
		http.Redirect(w, r, "/servers", 302)
	}
	fmt.Println(r.PostForm)

	p := &Page{HeaderMessage{Visible: "hidden"}, true, nil}
	//	p.Mess.Type = "danger"
	// p.Mess.Message = "Please fill out all fields"
	// p.Mess.Visible = ""

	if !isSession(r) {
		http.Redirect(w, r, "/login", 302)
	}
	s.Srvs, err = getServer()
	if err != nil {
		fmt.Println(err)
		p.Mess.Type = "Danger"
		p.Mess.Message = "Couldn't get Servers"
		p.Mess.Visible = ""
	}
	id := r.URL.Path[len("/server/"):]
	s.Curr, _ = strconv.Atoi(id)
	p.Info = s
	serverPage(w, r, p)
}

func addSrv(w http.ResponseWriter, r *http.Request) {
	var s Servers
	var err error

	p := &Page{HeaderMessage{Visible: "hidden"}, true, nil}
	if !isSession(r) {
		http.Redirect(w, r, "/login", 302)
	}
	r.ParseForm()
	if checkPost(r.PostForm, "name", "user", "pass", "addr", "port") {
		name := strings.TrimSpace(r.PostFormValue("name"))
		user := strings.TrimSpace(r.PostFormValue("user"))
		pass := strings.TrimSpace(r.PostFormValue("pass"))
		addr := strings.TrimSpace(r.PostFormValue("addr"))
		port := strings.TrimSpace(r.PostFormValue("port"))
		addServer(name, user, pass, port, addr)
		http.Redirect(w, r, "/servers", 302)
	}
	s.Srvs, err = getServer()
	if err != nil {
		fmt.Println(err)
		p.Mess.Type = "Danger"
		p.Mess.Message = "Couldn't get Servers"
		p.Mess.Visible = ""
	}
	p.Info = s
	addSrvPage(w, r, p)
}

// func delSrv(id int) {
//	id , _ = strconv.Atoi(id)
//	http.Redirect(w, r, "/servers", 302)
// }
