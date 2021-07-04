package main

import (
	"log"
	"net/http"
	"web-store/controller"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// 配置服务器结构
	server := http.Server{
		Addr:    "localhost:8081",
		Handler: nil,
	}

	// 路由器注册路由规则和处理器
	controller.RegisterRoutes()

	// 启动服务器
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
