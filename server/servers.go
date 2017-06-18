package main

type Servers struct {
        Srvs []Server
}

type Server struct {
	ID int
        Name string
        Username string
        Password string
        Port string
        Address string
        Status string
}

