package controller

import (
	"html/template"
	"log"
	"net/http"
	"web-store/model"
	"web-store/util"
)

func registerIndexRoutes() {
	http.HandleFunc("/index", Index)
}

func Index(w http.ResponseWriter, r *http.Request) {
	ok, sess := IsLogin(r)

	if !ok {
		sess = &model.Session{}
	}

	newProducts, err := model.GetIndexNewProducts()
	if err != nil {
		log.Println("从数据库获取首页新品产品发生错误:", err)
		http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
		return
	}

	hotProducts, err := model.GetIndexHotProducts()
	if err != nil {
		log.Println("从数据库获取首页热销良品发生错误:", err)
		http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
		return
	}

	recomProducts, err := model.GetIndexRecomProducts()
	if err != nil {
		log.Println("从数据库获取首页推荐产品发生错误:", err)
		http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
		return
	}

	indexPage := &model.IndexPage{
		NewProducts:   newProducts,
		HotProducts:   hotProducts,
		RecomProducts: recomProducts,
	}

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

	nav := &model.Nav{
		Categories:  categories,
		NavProducts: navProducts,
	}

	sess.IndexPage = indexPage
	sess.Nav = nav

	// 解析模板文件，并执行模板，生成包含动态数据的HTML文档，返回给浏览器
	t := template.New("layout")
	t, err = t.ParseFiles("./view/template/layout.html", "./view/template/index.html")
	if err != nil {
		log.Printf("解析模板文件发生错误：%v\n", err)
		http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
	} else {
		t.ExecuteTemplate(w, "layout", sess)
	}
}
