 package main

import (
	"net"
	"fmt"
	"bytes"
	"strconv"
	"io/ioutil"
	"encoding/json"
)

type config struct {
	Port string
	Username string
	Password string
}

var cfg config

// TODO error manage
func appendBytes(by ...[]byte) []byte {
	var ab bytes.Buffer

	for _, b := range by {
		ab.WriteString(strconv.Itoa(len(b)))
		ab.WriteString(":")
		ab.Write(b)
	}
	ab.WriteByte(0x03)
	return ab.Bytes()
}

func auth(conn net.Conn) {
	data := make([]byte, 512)
	conn.Read(data)
	login := explodeString(string(data))
	if login[0] == cfg.Username && login[1] == cfg.Password {
		fmt.Println("confirmed")
		StartMonitor(conn)
	} else {
		fmt.Println("wrong password")
	}
}

func main() {
	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	json.Unmarshal(b, &cfg)
	fmt.Println(cfg)
	ln, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("starting new comm")
			go auth(conn)
			// go StartMonitor(conn)
		}
	}
}
