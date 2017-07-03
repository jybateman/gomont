package main

import (
	"fmt"
	"net"
	"time"
	"strconv"
	"syscall"
	"strings"
	"os/exec"
	"io/ioutil"
	"encoding/json"
)

var timeout time.Duration = 0

type monitor struct {
	Name string
	Command string
	Frequence string
	Graph bool
	d time.Time
}

func (mon *monitor) execCmd(conn net.Conn) error {
	parts := strings.Fields(mon.Command)
	head := parts[0]
	parts = parts[1:len(parts)]
	cmd := exec.Command(head, parts...)
	out, err := cmd.Output()
	status := 0
	if err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if tmp, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				status = tmp.ExitStatus()
			}
		} else {
			fmt.Println(err)
			return err
		}
	}
	_, err = conn.Write(appendBytes([]byte(mon.Name), []byte(strconv.Itoa(status)), strconv.AppendBool(nil, mon.Graph), out))
	if err != nil {
		fmt.Println("terminating communication with server")
		conn.Close()
		return err
	}
	return nil
}

func setTimeout(mon []monitor) {
	for idx, m := range mon {
		t := m.d.Sub(time.Now())
		if idx == 0 || t < timeout {
			timeout = t
		}
	}
}

func updateTime(mon []monitor, conn net.Conn) error {
	t := time.Now()
	for idx, _ := range mon {
		st := mon[idx].d.Sub(time.Now())
		if st <= 0 {
			s, _ := time.ParseDuration(mon[idx].Frequence)
			mon[idx].d = t.Add(s)
			return mon[idx].execCmd(conn)
		}
	}
	return nil
}

func StartMonitor(conn net.Conn) {
	var mon []monitor
	b, err := ioutil.ReadFile("monitor.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	json.Unmarshal(b, &mon)
	t := time.Now()
	for idx, _ := range mon {
		s, _ := time.ParseDuration(mon[idx].Frequence)
		mon[idx].d = t.Add(s)
	}
	for {
		time.Sleep(timeout)
		err = updateTime(mon, conn)
		if err != nil {
			fmt.Println("closing thread")
			return
		}
		setTimeout(mon)
	}
}
