package main

import (
	"os"
	"fmt"
	"net"
	"time"
	"bytes"
	"strings"
	"strconv"
	"golang.org/x/net/websocket"
)

type Servers struct {
	Srvs []Server
	Curr int
}

type Server struct {
	ID  int
	Name string
	Username string
	Password string
	Port string
	Address string
	Status bool
	conn net.Conn
}

var addS chan Server
var delS chan int

func  (s *Server) readServer(ch chan string, quit chan bool) {
	data := make([]byte, 512)
	for {
		select {
		case <- quit:
			fmt.Println("terminating communication with server:", s.Name)
			s.conn.Close()
			return
		default:
			_, err := s.conn.Read(data)
			if err == nil {
				idx := bytes.LastIndexByte(data, 0x03)
				str := fmt.Sprintf("%d:%d%s", len(strconv.Itoa(s.ID)), s.ID, data[:idx])
				select {
				case ch <- str:
				case <- time.After(5 * time.Second):
					fmt.Println("terminating communication with server:", s.Name)
					s.conn.Close()
					s.conn = nil
					return
				}
			} else {
				fmt.Println("terminating communication with server:", s.Name)
				s.conn.Close()
				s.conn = nil
				return
			}
		}
	}
}

func (s *Server) storeData() {
	var f *os.File

	ch := make(chan string)
	quit := make(chan bool)
	os.Mkdir("./files/"+strconv.Itoa(s.ID), os.ModeDir | 0775)
	fmt.Println("Starting communication with:", s.Address+":"+s.Port)
	go s.readServer(ch, quit)
	for {
		select {
		case data := <- ch:
			res := explodeString(data)
			if len(res) > 4 {
				t := time.Now()
				_, err := os.Stat("./files/"+strconv.Itoa(s.ID)+"/"+res[1]+".csv")
				if !os.IsNotExist(err) {
					f, err = os.OpenFile("./files/"+strconv.Itoa(s.ID)+"/"+res[1]+".csv", os.O_APPEND | os.O_RDWR, os.ModeAppend)
				}  else {
					f, err = os.Create("./files/"+strconv.Itoa(s.ID)+"/"+res[1]+".csv")
					if err == nil {
						f.WriteString("Date,Value\n");
					}
				}
				if err == nil {
					_, err = f.WriteString(fmt.Sprintf("%d/%02d/%02d %02d:%02d:%02d,%s\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), res[4]))
					f.Sync()
					f.Close()
				}
			}
		}
	}
}

// TODO extend timeout
func (s *Servers) commJS(ws *websocket.Conn) {
	ch := make(chan string)
	quit := make(chan bool)
	for i, _ := range s.Srvs {
		if s.Srvs[i].conn != nil {
			fmt.Println("Starting communication with:", s.Srvs[i].Address+":"+s.Srvs[i].Port, i)
			go s.Srvs[i].readServer(ch, quit)
		}
	}
	for {
		select {
		case data := <- ch:
			_, err := ws.Write([]byte(data))
			if err != nil {
				fmt.Println("Terminating communication with websocket")
				fmt.Println(err)
				quit <- true
				ws.Close()
				return
			}
		case <- time.After(5 * time.Second):
			_, err := ws.Write(make([]byte, 0))
			if err != nil {
				fmt.Println("Terminating communication with websocket")
				fmt.Println(err)
				quit <- true
				ws.Close()
				return
			}
		}
	}
}

func dialWS(ws *websocket.Conn) {
	var err error
	var s Servers

	id := strings.Split(ws.Request().URL.String(), "/")
	s.Srvs, err = getServer()
	if len(id) > 2 {
		s.Curr, _ = strconv.Atoi(id[2])
	}
	for i, _ := range s.Srvs {
		if s.Srvs[i].ID == s.Curr {
			s.Srvs[i].conn, err = net.Dial("tcp", s.Srvs[i].Address+":"+s.Srvs[i].Port)
			s.Srvs[i].Status = true
			if err != nil {
				fmt.Println(err)
				s.Srvs[i].conn = nil
				s.Srvs[i].Status = false
			}
		}
	}
	s.commJS(ws)
}

func dialServer() {
	var err error
	var s Servers

	s.Srvs, err = getServer()
	for i, _ := range s.Srvs {
		if s.Srvs[i].conn == nil {
			s.Srvs[i].conn, err = net.Dial("tcp", s.Srvs[i].Address+":"+s.Srvs[i].Port)
			if err == nil && s.Srvs[i].conn != nil {
				fmt.Println("Connected to server")
				go s.Srvs[i].storeData()
			} else {
				fmt.Println("Couldn't connect to server")
			}
		}
	}

	select {
	case srv := <- addS:
		s.Srvs = append(s.Srvs, srv)
		break
	case id := <- delS:
		_ = id
		break
	}
}
