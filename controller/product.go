package controller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"web-store/model"
	"web-store/util"
)

func registerProductRoutes() {
	// http.HandleFunc("/products", getProducts)
	http.HandleFunc("/products", getPageProductsByPrice)
	// http.HandleFunc("/productsByPrice", getPageProductsByPrice)
	http.HandleFunc("/product", getProduct)
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
	ok, sess := IsLogin(r)
	if ok {
		// 设置分页结构
		page.IsLogin = true
		page.Username = sess.Username
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

	// 获取产品类别id
	category_id := r.FormValue("category_id")

	var page *model.PageProduct
	var err error

	if maxPrice == "" {
		if category_id == "" {
			// 从数据库获取当前产品列表分页结构
			page, err = model.GetPageProducts(pageNo)
			if err != nil {
				log.Printf("查询数据库发生错误：%v", err)
				http.Error(w, "服务器内部发生错误", http.StatusInternalServerError)
				return
			}
		} else {
			page, err = model.GetPageProductsByCategoryID(pageNo, category_id)
			if err != nil {
				log.Printf("查询数据库发生错误：%v", err)
				http.Error(w, "服务器内部发生错误", http.StatusInternalServerError)
				return
			}
			// 将产品类别id设置到分页结构，让产品列表模板使用
			page.Category_id = category_id
		}

	} else {
		if category_id == "" {
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
		} else {
			page, err = model.GetPageProductsByPriceAndCategoryID(pageNo, category_id, minPrice, maxPrice)
			if err != nil {
				log.Printf("查询数据库发生错误：%v", err)
				http.Error(w, "服务器内部发生错误", http.StatusInternalServerError)
				return
			}
			// 将产品类别id和价格区间设置到分页结构，让产品列表模板使用
			page.Category_id = category_id
			page.MinPrice = minPrice
			page.MaxPrice = maxPrice
		}
	}

	// 检查用户是否已经登录
	ok, sess := IsLogin(r)
	if ok {
		// 设置分页结构
		page.IsLogin = true
		page.Username = sess.Username
	}

	cates, err := model.GetCategories()
	if err != nil {
		log.Println(err)
		http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
		return
	}
	page.Categories = cates

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

func getProduct(w http.ResponseWriter, r *http.Request) {
	// 从查询字符串参数获取产品id
	pid := r.FormValue("pid")

	// 从数据库获取产品
	p, err := model.GetProduct(pid)

	if err == model.ErrNotFoundProduct {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err != nil {
		log.Println(err)
		http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
		return
	}

	strCategory_id := strconv.Itoa(p.Category_id)
	cate, err := model.GetCategory(strCategory_id)
	if err != nil {
		log.Println(err)
		http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
		return
	}
	p.Category = cate

	// 判断会员是否已经登录
	ok, sess := IsLogin(r)
	if ok {
		// 产品结构设置模板需要使用的字段
		p.IsLogin = true
		p.Username = sess.Username
	}

	// 解析模板文件，执行模板结合动态数据，生成最终HTML文档，传递给ResponseWriter响应客户端
	t, err := template.ParseFiles("./view/template/layout.html", "./view/template/product.html")
	if err != nil {
		log.Printf("解析模板文件发生错误：%v\n", err)
		http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
	} else {
		t.ExecuteTemplate(w, "layout", p)
	}
}
