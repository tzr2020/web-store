package controller

import "net/http"

func RegsiRoutes() {
	// 给静态资源注册路由
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("view/static"))))

	// 后台管理页面资源注册路由
	http.Handle("/manage/", http.StripPrefix("/manage/", http.FileServer(http.Dir("view/template/manage"))))

	// 前台页面处理器注册路由
	registerHomeRoutes()
	registerAboutRoutes()
	registerUserRoutes()
	registerProductRoutes()
	regsiterCartRoutes()
	registerOrderRoutes()
	registerIndexRoutes()

	// API处理器注册路由
	registerAPIUserRoutes()
}
