package main

import (
	"golang.org/x/net/websocket"
	"net/http"
)

func main() {
	go dialServer()
	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/servers", servers)
	http.HandleFunc("/server/", server)
	http.HandleFunc("/addserver", addSrv)
	http.Handle("/ws/", websocket.Handler(dialWS))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("fonts"))))
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("files"))))
	http.ListenAndServe(":9000", nil)
}
