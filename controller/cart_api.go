package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"web-store/model"
	"web-store/util"
)

func registerAPICartRouters() {
	http.HandleFunc("/api/carts", carts)
	http.HandleFunc("/api/cart", cart)
}

func carts(w http.ResponseWriter, r *http.Request) {
	// 从查询字符串参数获取数据
	pageNo := r.FormValue("pageNo")
	pageSize := r.FormValue("pageSize")
	// 数据类型转换
	intPageNo, err := strconv.Atoi(pageNo)
	intPageSize, err := strconv.Atoi(pageSize)
	if err != nil {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}

	// 查询数据库，获取购物车列表
	list, err := model.GetCartsPage(intPageNo, intPageSize)
	if err != nil {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}
	// 查询数据库，获取购物车表的记录条数
	count, err := util.GetJsonDataCount("carts")
	if err != nil {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}

	// 返回代表成功的JSON响应
	util.ResponseWriteJson(w, util.Json{
		Code:  200,
		Msg:   "获取购物车列表成功",
		Count: count,
		Data:  list,
	})
}

func cart(w http.ResponseWriter, r *http.Request) {
	// 解码包体JSON数据到结构
	dec := json.NewDecoder(r.Body)
	c := model.Cart{}
	err := dec.Decode(&c)
	if err != nil {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}

	switch r.Method {
	case http.MethodPost:
		// 数据库添加购物车
		if err = c.Add(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库添加购物车成功",
		})
	case http.MethodPut:
		// 数据库更新购物车
		if err = c.Update(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库更新购物车成功",
		})
	case http.MethodDelete:
		// 数据库删除购物车
		if err = c.Delete(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库删除购物车成功",
		})
	}
}
