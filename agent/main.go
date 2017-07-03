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

func main() {
	var cfg config

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
			go StartMonitor(conn)
		}
	}
}
