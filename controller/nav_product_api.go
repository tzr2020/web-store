package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"web-store/model"
	"web-store/util"
)

func registerAPINavProductRouters() {
	http.HandleFunc("/api/nav-products", navProducts)
	http.HandleFunc("/api/nav-product", navProduct)
}

// navProducts 获取导航栏产品列表，根据当前页的页码和每页记录条数
func navProducts(w http.ResponseWriter, r *http.Request) {
	// 从查询字符串获取数据
	pageNo := r.FormValue("pageNo")     // 当前页页码
	pageSize := r.FormValue("pageSize") // 每页记录条数
	// 数据类型转换
	intPageNo, err := strconv.Atoi(pageNo)
	intPageSize, err := strconv.Atoi(pageSize)
	if err != nil {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}

	switch r.Method {
	case http.MethodGet:
		// 查询数据库，获取列表的总记录条数
		count, err := util.GetJsonDataCount("nav_products")
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		// 查询数据库，获取列表，根据当前页的页码和每页记录条数
		list, err := model.GetNavProductPage(intPageNo, intPageSize)
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		// 将JSON结构编码为JSON格式数据后写入响应
		j := util.Json{
			Code:  200,
			Msg:   "获取导航栏产品列表成功",
			Count: count,
			Data:  list,
		}
		util.ResponseWriteJson(w, j)
	}
}

func navProduct(w http.ResponseWriter, r *http.Request) {
	// 解码包体JSON数据到结构
	dec := json.NewDecoder(r.Body)
	np := model.NavProduct{}
	err := dec.Decode(&np)
	if err != nil {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}

	switch r.Method {
	case http.MethodPost:
		// 添加导航栏产品

		// 数据库添加导航栏产品
		if err = np.Add(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库添加导航栏产品成功",
		})
	case http.MethodPut:
		// 更新导航栏产品

		// 数据库更新导航栏产品
		if err = np.Update(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库更新导航栏产品成功",
		})
	case http.MethodDelete:
		// 删除导航栏产品

		// 数据库删除导航栏产品
		if err = np.Delete(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库删除导航栏产品成功",
		})
	}
}
