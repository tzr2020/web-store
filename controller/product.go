package controller

import (
	"html/template"
	"log"
	"net/http"
	"web-store/model"
)

func registerProductRoutes() {
	http.HandleFunc("/products", getProducts)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	ps, err := model.GetProducts()

	if err != nil {
		log.Printf("从数据库获取数据发生错误：%v", err)
		http.Error(w, "服务器内部发生错误", http.StatusInternalServerError)
		return
	}

	t, err := template.ParseFiles("./view/page/product/products.html")

	if err != nil {
		log.Printf("解析模板文件发生错误：%v", err)
		http.Error(w, "服务器内部发生错误", http.StatusInternalServerError)
	} else {
		t.Execute(w, ps)
	}
}
