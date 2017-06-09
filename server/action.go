package main

import (
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	loginPage(w, r)
}

func signup(w http.ResponseWriter, r *http.Request) {
	signupPage(w, r)
}

func logout(w http.ResponseWriter, r *http.Request) {

}

func servers(w http.ResponseWriter, r *http.Request) {

}
