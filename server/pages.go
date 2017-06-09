package main

import (
	"fmt"
	"net/http"
	"html/template"
)

func loginPage(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("html/login.html", "html/header.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = r.ParseForm()
	if err != nil {
		return
	}
	tpl.Execute(w, nil)
}

func signupPage(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("html/signup.html", "html/header.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = r.ParseForm()
	if err != nil {
		return
	}
	tpl.Execute(w, nil)
}

func serversPage() {

}
