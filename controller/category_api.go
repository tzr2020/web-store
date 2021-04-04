package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"web-store/model"
	"web-store/util"
)

func registerAPICategoryRouters() {
	http.HandleFunc("/api/categories", categories)
	http.HandleFunc("/api/category", category)
}

func categories(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// 获取产品类别列表

		// 查询数据库，获取产品类别列表
		categories, err := model.GetCategories()
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		// 返回成功的JSON数据
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "获取产品类别列表成功",
			Data: categories,
		})
	}
}

func category(w http.ResponseWriter, r *http.Request) {
	// 解码包体JSON数据到结构
	dec := json.NewDecoder(r.Body)
	cate := model.Category{}
	err := dec.Decode(&cate)
	if err != nil {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}

	switch r.Method {
	case http.MethodPost:
		// 添加产品类别

		// 数据库添加产品类别
		if err = cate.Add(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库添加产品类别成功",
		})

	case http.MethodPut:
		// 更新产品类别

		// 数据库更新产品类别
		if err = cate.Update(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库更新产品类别成功",
		})

	case http.MethodDelete:
		// 删除产品类别

		// 数据库删除产品类别
		if err = cate.Delete(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库删除产品类别成功",
		})
	}
}
