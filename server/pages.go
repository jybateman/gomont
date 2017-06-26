package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type Page struct {
	Mess HeaderMessage
	Nav bool
	Info interface{}
}

type HeaderMessage struct {
	Visible string
	Type string
	Message string
}

func loginPage(w http.ResponseWriter, r *http.Request, p *Page) {
	tpl, err := template.ParseFiles("html/login.html", "html/header.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}
	tpl.Execute(w, p)
}

func signupPage(w http.ResponseWriter, r *http.Request, p *Page) {
	tpl, err := template.ParseFiles("html/signup.html", "html/header.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}
	tpl.Execute(w, p)
}

func serversPage(w http.ResponseWriter, r *http.Request, p *Page) {
	tpl, err := template.ParseFiles("html/home.html", "html/header.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}
	tpl.Execute(w, p)
}

func serverPage(w http.ResponseWriter, r *http.Request, p *Page) {
	tpl, err := template.ParseFiles("html/server.html", "html/header.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}
	tpl.Execute(w, p)
}

func addSrvPage(w http.ResponseWriter, r *http.Request, p *Page) {
	tpl, err := template.ParseFiles("html/addserver.html", "html/header.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}
	tpl.Execute(w, p)
}
