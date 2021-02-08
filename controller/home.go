package controller

import (
	"html/template"
	"log"
	"net/http"
)

func registerHomeRoutes() {
	http.HandleFunc("/home", handlerHome)
}

func handlerHome(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./view/template/layout.html", "./view/template/home.html")
	if err != nil {
		log.Fatalln(err)
	}
	err = t.ExecuteTemplate(w, "layout", "Hello world!")
	if err != nil {
		log.Fatalln(err)
	}
}
