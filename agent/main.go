 package main

import (
	"net"
	"fmt"
	"bytes"
	"strconv"
)

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
	ln, err := net.Listen("tcp", ":4242")
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
