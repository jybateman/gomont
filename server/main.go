package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/servers", servers)
	http.HandleFunc("/addserver", addSrv)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
        http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
        http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("files"))))
	http.ListenAndServe(":9000", nil)
}
