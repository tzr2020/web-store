package controller

import "net/http"

func RegsiRoutes() {
	// 给静态资源注册路由
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("view/static"))))

	// 给各个处理器注册路由
	registerHomeRoutes()
	registerAboutRoutes()
	registerUserRoutes()
	registerProductRoutes()
	regsiterCartRoutes()
	registerOrderRoutes()
}
