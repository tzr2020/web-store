package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"web-store/controller"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

const (
	username = "root"
	password = "123456"
	ip       = "localhost"
	port     = "3306"
	database = "store"
)

func init() {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		username, password, ip, port, database)
	Db, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalln(err)
	}

	err = Db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Database connected!")
}

func main() {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}

	controller.RegsiRoutes()

	err = server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
