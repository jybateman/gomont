package main

import (
	"fmt"
	
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func addServer(name, user, pass, port, addr string) error {
	db, err := sql.Open("mysql",
		"root:helloworld@tcp(127.0.0.1:3306)/gomont")
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO server (name, username, password, address, port) VALUE (?, ?, ?, ?, ?)",
		name, user, pass, addr, port)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func getServer() ([]Server, error) {
	var svrs []Server
	
	db, err := sql.Open("mysql",
		"root:helloworld@tcp(127.0.0.1:3306)/gomont")
	if err != nil {
		return svrs, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT id, name, username, password, address, port FROM server")
	if err != nil {
		return svrs, err
	}
	for rows.Next() {
		var svr Server
		rows.Scan(&svr.ID, &svr.Name, &svr.Username, &svr.Password, &svr.Address, &svr.Port)
		svrs = append(svrs, svr)
	}
	defer rows.Close()
	return svrs, nil
}

func checkAccount(user, pass string) bool {
	var res int
	
	db, err := sql.Open("mysql",
		"root:helloworld@tcp(127.0.0.1:3306)/gomont")
	if err != nil {
		return false
	}
	defer db.Close()
	rows, err := db.Query("SELECT COUNT(*) FROM admin WHERE username=? AND password=?",
		user, pass)
	if err != nil {
		return false
	}
	defer rows.Close()
	rows.Next()
	err = rows.Scan(&res)
	if err != nil || res == 0 {
		return false
	}
	return true
}

func addAdmin(user, pass string) error {
	db, err := sql.Open("mysql",
		"root:helloworld@tcp(127.0.0.1:3306)/gomont")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO admin VALUE (?, ?)",
		user, pass)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func isAdmin() bool {
	var res int
	
	db, err := sql.Open("mysql",
		"root:helloworld@tcp(127.0.0.1:3306)/gomont")
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer db.Close()
	rows, err := db.Query("SELECT COUNT(*) FROM admin")
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer rows.Close()
	rows.Next()
	err = rows.Scan(&res)
	if err != nil || res == 0 {
		return false
	}
	return true
}

