package controller

import "net/http"

func RegsiRoutes() {
	// 给静态资源注册路由
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("view/static"))))

	registerHomeRoutes()
	registerAboutRoutes()
	registerUserRoutes()
	registerProductRoutes()
}
