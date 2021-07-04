package controller

import "net/http"

func RegisterRoutes() {
	// 静态文件注册路由
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("view/static"))))

	// 后台管理页面注册路由
	http.Handle("/manage/", http.StripPrefix("/manage/", http.FileServer(http.Dir("view/template/manage"))))
	registerManageRoutes()

	// 前台处理器注册路由
	registerUserRoutes()
	registerProductRoutes()
	regsiterCartRoutes()
	registerOrderRoutes()
	registerIndexRoutes()

	// API处理器注册路由（处理异步请求）
	registerAPIUserRoutes()
	registerAPIProductRouters()
	registerAPICategoryRouters()
	registerAPIAdminRouters()
	registerAPIUserAddressRouters()
	registerAPICartRouters()
	registerAPICartitemRouters()
	registerAPIIndexProductRouters()
	registerAPINavProductRouters()
	registerAPIOrderRouters()
	registerAPIUserSessionRouters()
	registerAPIAdminSessionRouters()
}
