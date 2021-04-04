package util

import (
	"html/template"
	"log"
	"net/http"
)

// ExecuteTpl 解析一个模板文件，执行模板，返回响应页面
func ExecuteTpl(w http.ResponseWriter, filename string) {
	t, err := template.ParseFiles(filename)
	if err != nil {
		log.Println(err)
		http.Error(w, ErrServerInside.Error(), 500)
		return
	}
	if err = t.Execute(w, nil); err != nil {
		log.Println(err)
		http.Error(w, ErrServerInside.Error(), 500)
	}
}

// ExecuteTowTpl 解析两个模板文件，第一个是layout，第二个是content，执行模板，返回响应页面
func ExecuteTowTpl(w http.ResponseWriter, filenames [2]string) {
	t := template.New("layout")
	t, err = t.ParseFiles(filenames[0], filenames[1])
	if err != nil {
		log.Println(err)
		http.Error(w, ErrServerInside.Error(), 500)
		return
	}
	if err = t.ExecuteTemplate(w, "layout", nil); err != nil {
		log.Println(err)
		http.Error(w, ErrServerInside.Error(), 500)
	}
}
