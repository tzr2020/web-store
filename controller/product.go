package controller

import (
	"html/template"
	"log"
	"net/http"
	"web-store/model"
)

func registerProductRoutes() {
	// http.HandleFunc("/products", getProducts)
	http.HandleFunc("/products", getPageProductsByPrice)
	// http.HandleFunc("/productsByPrice", getPageProductsByPrice)
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

func getPageProducts(w http.ResponseWriter, r *http.Request) {
	// 获取页码
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}

	// 从数据库获取当前产品列表分页结构
	page, err := model.GetPageProducts(pageNo)
	if err != nil {
		log.Printf("查询数据库发生错误：%v", err)
		http.Error(w, "服务器内部发生错误", http.StatusInternalServerError)
		return
	}

	// 检查用户是否已经登录
	ok, username := IsLogin(r)
	if ok {
		// 设置分页结构
		page.IsLogin = true
		page.Username = username
	}

	// 解析模板文件
	t, err := template.ParseFiles("./view/template/layout.html", "./view/template/products.html")
	// t, err := template.ParseFiles("./view/page/product/products.html")

	// 执行模板，生成HTML文档，返回页面
	if err != nil {
		log.Printf("解析模板文件发生错误：%v", err)
		http.Error(w, "服务器内部发生错误", http.StatusInternalServerError)
	} else {
		t.ExecuteTemplate(w, "layout", page)
		// t.Execute(w, page)
	}
}

func getPageProductsByPrice(w http.ResponseWriter, r *http.Request) {
	// 获取页码
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}

	// 获取价格区间
	minPrice := r.FormValue("min")
	maxPrice := r.FormValue("max")
	if minPrice == "" {
		minPrice = "0"
	}

	var page *model.PageProduct
	var err error

	if maxPrice == "" {
		// 从数据库获取当前产品列表分页结构
		page, err = model.GetPageProducts(pageNo)
		if err != nil {
			log.Printf("查询数据库发生错误：%v", err)
			http.Error(w, "服务器内部发生错误", http.StatusInternalServerError)
			return
		}
	} else {
		// 从数据库获取当前产品列表分页结构，根据价格
		page, err = model.GetPageProductsByPrice(pageNo, minPrice, maxPrice)
		if err != nil {
			log.Printf("查询数据库发生错误：%v", err)
			http.Error(w, "服务器内部发生错误", http.StatusInternalServerError)
			return
		}
		// 将价格区间设置到分页结构，让产品列表模板使用
		page.MinPrice = minPrice
		page.MaxPrice = maxPrice
	}

	// 检查用户是否已经登录
	ok, username := IsLogin(r)
	if ok {
		// 设置分页结构
		page.IsLogin = true
		page.Username = username
	}

	// 解析模板文件
	t, err := template.ParseFiles("./view/template/layout.html", "./view/template/products.html")
	// t, err := template.ParseFiles("./view/page/product/products.html")

	// 执行模板，生成HTML文档，返回页面
	if err != nil {
		log.Printf("解析模板文件发生错误：%v", err)
		http.Error(w, "服务器内部发生错误", http.StatusInternalServerError)
	} else {
		t.ExecuteTemplate(w, "layout", page)
		// t.Execute(w, page)
	}
}
