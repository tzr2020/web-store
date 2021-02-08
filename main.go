package main

import (
	"log"
	"net/http"
	"web-store/controller"
)

func main() {
	server := http.Server{
		Addr:    "localhost:8081",
		Handler: nil,
	}

	controller.RegsiRoutes()

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
