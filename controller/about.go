package controller

import (
	"html/template"
	"log"
	"net/http"
	"regexp"
)

func registerAboutRoutes() {
	http.HandleFunc("/about", handlerAbout)
	http.HandleFunc("/about/", handlerAbout2)
}

func handlerAbout(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./view/template/layout.html", "./view/template/about.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = t.ExecuteTemplate(w, "layout", "About me.")
	if err != nil {
		log.Fatalln(err)
	}
}

func handlerAbout2(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./view/template/layout.html", "./view/template/about.html")
	if err != nil {
		log.Fatalln(err)
	}

	pattern, err := regexp.Compile(`/about/(\d+)`) // 示例：/about/123
	matches := pattern.FindStringSubmatch(r.URL.Path)
	if len(matches) > 0 {
		err = t.ExecuteTemplate(w, "layout", "About "+matches[1]+".")
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		// 404
		w.WriteHeader(http.StatusNotFound)
	}
}
