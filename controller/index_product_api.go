package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"web-store/model"
	"web-store/util"
)

func registerAPIIndexProductRouters() {
	http.HandleFunc("/api/index/hot-products", indexHotProducts)
	http.HandleFunc("/api/index/hot-product", indexHotProduct)
	http.HandleFunc("/api/index/new-products", indexNewProducts)
	http.HandleFunc("/api/index/new-product", indexNewProduct)
	http.HandleFunc("/api/index/recom-products", indexRecomProducts)
	http.HandleFunc("/api/index/recom-product", indexRecomProduct)
}

// indexHotProducts 获取首页热卖产品列表，根据当前页的页码和每页记录条数
func indexHotProducts(w http.ResponseWriter, r *http.Request) {
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
		count, err := util.GetJsonDataCount("index_hot_products")
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		// 查询数据库，获取列表，根据当前页的页码和每页记录条数
		list, err := model.GetIndexHotProductPage(intPageNo, intPageSize)
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		// 将JSON结构编码为JSON格式数据后写入响应
		j := util.Json{
			Code:  200,
			Msg:   "获取首页热卖产品列表成功",
			Count: count,
			Data:  list,
		}
		util.ResponseWriteJson(w, j)
	}
}

func indexHotProduct(w http.ResponseWriter, r *http.Request) {
	// 解码包体JSON数据到结构
	dec := json.NewDecoder(r.Body)
	ihp := model.IndexHotProduct{}
	err := dec.Decode(&ihp)
	if err != nil {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}

	switch r.Method {
	case http.MethodPost:
		// 添加首页热卖产品

		// 数据库添加首页热卖产品
		if err = ihp.Add(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库添加首页热卖产品成功",
		})
	case http.MethodPut:
		// 更新首页热卖产品

		// 数据库更新首页热卖产品
		if err = ihp.Update(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库更新首页热卖产品成功",
		})
	case http.MethodDelete:
		// 删除首页热卖产品

		// 数据库删除首页热卖产品
		if err = ihp.Delete(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库删除首页热卖产品成功",
		})
	}
}

// indexNewProducts 获取首页最新产品列表，根据当前页的页码和每页记录条数
func indexNewProducts(w http.ResponseWriter, r *http.Request) {
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
		count, err := util.GetJsonDataCount("index_new_products")
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		// 查询数据库，获取列表，根据当前页的页码和每页记录条数
		list, err := model.GetIndexNewProductPage(intPageNo, intPageSize)
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		// 将JSON结构编码为JSON格式数据后写入响应
		j := util.Json{
			Code:  200,
			Msg:   "获取首页最新产品列表成功",
			Count: count,
			Data:  list,
		}
		util.ResponseWriteJson(w, j)
	}
}

func indexNewProduct(w http.ResponseWriter, r *http.Request) {
	// 解码包体JSON数据到结构
	dec := json.NewDecoder(r.Body)
	inp := model.IndexNewProduct{}
	err := dec.Decode(&inp)
	if err != nil {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}

	switch r.Method {
	case http.MethodPost:
		// 添加首页最新产品

		// 数据库添加首页最新产品
		if err = inp.Add(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库添加首页最新产品成功",
		})
	case http.MethodPut:
		// 更新首页最新产品

		// 数据库更新首页最新产品
		if err = inp.Update(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库更新首页最新产品成功",
		})
	case http.MethodDelete:
		// 删除首页最新产品

		// 数据库删除首页最新产品
		if err = inp.Delete(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库删除首页最新产品成功",
		})
	}
}

// indexRecomProducts 获取首页推荐产品列表，根据当前页的页码和每页记录条数
func indexRecomProducts(w http.ResponseWriter, r *http.Request) {
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
		count, err := util.GetJsonDataCount("index_recom_products")
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		// 查询数据库，获取列表，根据当前页的页码和每页记录条数
		list, err := model.GetIndexRecomProductPage(intPageNo, intPageSize)
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		// 将JSON结构编码为JSON格式数据后写入响应
		j := util.Json{
			Code:  200,
			Msg:   "获取首页推荐产品列表成功",
			Count: count,
			Data:  list,
		}
		util.ResponseWriteJson(w, j)
	}
}

func indexRecomProduct(w http.ResponseWriter, r *http.Request) {
	// 解码包体JSON数据到结构
	dec := json.NewDecoder(r.Body)
	irp := model.IndexRecomProduct{}
	err := dec.Decode(&irp)
	if err != nil {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}

	switch r.Method {
	case http.MethodPost:
		// 添加首页推荐产品

		// 数据库添加首页推荐产品
		if err = irp.Add(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库添加首页推荐产品成功",
		})
	case http.MethodPut:
		// 更新首页推荐产品

		// 数据库更新首页推荐产品
		if err = irp.Update(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库更新首页推荐产品成功",
		})
	case http.MethodDelete:
		// 删除首页推荐产品

		// 数据库删除首页推荐产品
		if err = irp.Delete(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库删除首页推荐产品成功",
		})
	}
}
