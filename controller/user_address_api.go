package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"web-store/model"
	"web-store/util"
)

func registerAPIUserAddressRouters() {
	http.HandleFunc("/api/user/addresses", userAddresses)
	http.HandleFunc("/api/user/address", userAddress)
}

// userAddresses 返回用户地址列表的JSON响应
func userAddresses(w http.ResponseWriter, r *http.Request) {
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

	// 查询数据库，获取用户地址列表
	as, err := model.GetAddressPage(intPageNo, intPageSize)
	if err != nil {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}
	// 查询数据库，获取用户地址表的记录条数
	count, err := util.GetJsonDataCount("addresses")
	if err != nil {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}

	// 返回代表成功的JSON响应
	util.ResponseWriteJson(w, util.Json{
		Code:  200,
		Msg:   "获取用户地址列表成功",
		Count: count,
		Data:  as,
	})
}

func userAddress(w http.ResponseWriter, r *http.Request) {
	// 解码包体JSON数据到结构
	dec := json.NewDecoder(r.Body)
	a := model.Address{}
	err := dec.Decode(&a)
	if err != nil {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}

	switch r.Method {
	case http.MethodPost:
		// 数据库添加用户地址
		if err = a.Add(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库添加用户地址成功",
		})
	case http.MethodPut:
		// 数据库更新用户地址
		if err = a.Update(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库更新用户地址成功",
		})
	case http.MethodDelete:
		// 数据库删除用户地址
		if err = a.Delete(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库删除用户地址成功",
		})
	}
}
