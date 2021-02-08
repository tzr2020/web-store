package main

import (
	"log"
	"net/http"
	"web-store/controller"
)

func main() {
	// 配置服务器结构
	server := http.Server{
		Addr:    "localhost:8081",
		Handler: nil,
	}

	// 路由器注册路由规则和处理器
	controller.RegsiRoutes()

	// 启动服务器
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
