package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"

	"golang.org/x/net/websocket"
)

type config struct {
	Port string
	Mysql sqlConf
}

var conf config

func main() {
	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	json.Unmarshal(b, &conf)

	go dialServer()
	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/servers", servers)
	http.HandleFunc("/server/", server)
	http.HandleFunc("/addserver", addSrv)
	http.HandleFunc("/editserver/", editSrv)
	http.Handle("/ws/", websocket.Handler(dialWS))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("fonts"))))
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("files"))))
	http.ListenAndServe(":"+conf.Port, nil)
}
