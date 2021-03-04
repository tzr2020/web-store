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

// func getProducts(w http.ResponseWriter, r *http.Request) {
// 	ps, err := model.GetProducts()

// 	if err != nil {
// 		log.Printf("从数据库获取数据发生错误：%v", err)
// 		http.Error(w, "服务器内部发生错误", http.StatusInternalServerError)
// 		return
// 	}

// 	t, err := template.ParseFiles("./view/page/product/products.html")

// 	if err != nil {
// 		log.Printf("解析模板文件发生错误：%v", err)
// 		http.Error(w, "服务器内部发生错误", http.StatusInternalServerError)
// 	} else {
// 		t.Execute(w, ps)
// 	}
// }

// getPageProducts 返回分页的产品列表页面
// func getPageProducts(w http.ResponseWriter, r *http.Request) {
// 	// 获取页码
// 	pageNo := r.FormValue("pageNo")
// 	if pageNo == "" {
// 		pageNo = "1"
// 	}

// 	// 从数据库获取当前产品列表分页结构
// 	page, err := model.GetPageProducts(pageNo)
// 	if err != nil {
// 		log.Printf("查询数据库发生错误：%v", err)
// 		http.Error(w, "服务器内部发生错误", http.StatusInternalServerError)
// 		return
// 	}

// 	// 检查用户是否已经登录
// 	ok, sess := IsLogin(r)
// 	if ok {
// 		// 设置分页结构
// 		page.IsLogin = true
// 		page.Username = sess.Username
// 	}

// 	// 解析模板文件
// 	t, err := template.ParseFiles("./view/template/layout.html", "./view/template/products.html")
// 	// t, err := template.ParseFiles("./view/page/product/products.html")

// 	// 执行模板，生成HTML文档，返回页面
// 	if err != nil {
// 		log.Printf("解析模板文件发生错误：%v", err)
// 		http.Error(w, "服务器内部发生错误", http.StatusInternalServerError)
// 	} else {
// 		t.ExecuteTemplate(w, "layout", page)
// 		// t.Execute(w, page)
// 	}
// }

// getPageProductsByPrice 根据请求参数（价格区间，产品类别）来查询数据库得到相应的产品列表分页结构后，返回产品列表页面
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

	cates, err := model.GetCategories()
	if err != nil {
		log.Println(err)
		http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
		return
	}
	page.Categories = cates

	categories, err := model.GetCategories()
	if err != nil {
		log.Println("从数据库获取所有产品类别发生错误:", err)
		http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
		return
	}

	navProducts, err := model.GetNavProducts()
	if err != nil {
		log.Println("从数据库获取导航栏产品类别发生错误:", err)
		http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
		return
	}

	// 检查用户是否已经登录
	ok, sess := IsLogin(r)
	if !ok {
		sess := &model.Session{}
		sess.PageProduct = page
		sess.Nav = &model.Nav{
			Categories:  categories,
			NavProducts: navProducts,
		}
		// 解析模板文件
		t, err := template.ParseFiles("./view/template/layout.html", "./view/template/products.html")
		// 执行模板，生成HTML文档，返回页面
		if err != nil {
			log.Printf("解析模板文件发生错误：%v", err)
			http.Error(w, "服务器内部发生错误", http.StatusInternalServerError)
		} else {
			t.ExecuteTemplate(w, "layout", sess)
			// t.Execute(w, page)
		}
	} else {
		sess.PageProduct = page
		sess.Nav = &model.Nav{
			Categories:  categories,
			NavProducts: navProducts,
		}
	}

	// 解析模板文件
	t, err := template.ParseFiles("./view/template/layout.html", "./view/template/products.html")
	// t, err := template.ParseFiles("./view/page/product/products.html")

	// 执行模板，生成HTML文档，返回页面
	if err != nil {
		log.Printf("解析模板文件发生错误：%v", err)
		http.Error(w, "服务器内部发生错误", http.StatusInternalServerError)
	} else {
		t.ExecuteTemplate(w, "layout", sess)
		// t.Execute(w, page)
	}
}

// getProduct 根据请求参数（产品id），返回产品详情页面
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

	categories, err := model.GetCategories()
	if err != nil {
		log.Println("从数据库获取所有产品类别发生错误:", err)
		http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
		return
	}

	navProducts, err := model.GetNavProducts()
	if err != nil {
		log.Println("从数据库获取导航栏产品类别发生错误:", err)
		http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
		return
	}

	// 判断会员是否已经登录
	ok, sess := IsLogin(r)
	if !ok {
		sess := &model.Session{}
		sess.Product = p
		sess.Nav = &model.Nav{
			Categories:  categories,
			NavProducts: navProducts,
		}
		// 解析模板文件，执行模板结合动态数据，生成最终HTML文档，传递给ResponseWriter响应客户端
		t, err := template.ParseFiles("./view/template/layout.html", "./view/template/product.html")
		if err != nil {
			log.Printf("解析模板文件发生错误：%v\n", err)
			http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
		} else {
			t.ExecuteTemplate(w, "layout", sess)
		}
	} else {
		sess.Product = p
		sess.Nav = &model.Nav{
			Categories:  categories,
			NavProducts: navProducts,
		}
	}

	// 解析模板文件，执行模板结合动态数据，生成最终HTML文档，传递给ResponseWriter响应客户端
	t, err := template.ParseFiles("./view/template/layout.html", "./view/template/product.html")
	if err != nil {
		log.Printf("解析模板文件发生错误：%v\n", err)
		http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
	} else {
		t.ExecuteTemplate(w, "layout", sess)
	}
}
