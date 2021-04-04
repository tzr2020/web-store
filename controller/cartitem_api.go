package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"web-store/model"
	"web-store/util"
)

func registerAPICartitemRouters() {
	http.HandleFunc("/api/cartitems", cartitems)
	http.HandleFunc("/api/cartitem", cartitem)
}

func cartitems(w http.ResponseWriter, r *http.Request) {
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
	list, err := model.GetCartitemsPage(intPageNo, intPageSize)
	if err != nil {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}
	// 查询数据库，获取购物车表的记录条数
	count, err := util.GetJsonDataCount("cart_items")
	if err != nil {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}

	// 返回代表成功的JSON响应
	util.ResponseWriteJson(w, util.Json{
		Code:  200,
		Msg:   "获取购物项列表成功",
		Count: count,
		Data:  list,
	})
}

func cartitem(w http.ResponseWriter, r *http.Request) {
	// 解码包体JSON数据到结构
	dec := json.NewDecoder(r.Body)
	cit := model.CartItem{}
	err := dec.Decode(&cit)
	if err != nil {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}

	switch r.Method {
	case http.MethodPost:
		// 添加购物项

		// 维护购物项结构体的产品字段
		strProductID := strconv.Itoa(cit.ProductID)
		product, err := model.GetProduct(strProductID)
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		cit.Product = product

		// 数据库添加购物项，分两种情况
		citem, err := model.GetCartItemByCartIDAndProductID(cit.CartID, strProductID)
		if citem == nil {
			// 购物车还没存在该产品的购物项，添加新的购物项
			err = model.AddCartItem(&cit)
			if err != nil {
				log.Println(err)
				util.ResponseWriteJsonOfInsideServer(w)
				return
			}
			// 数据库更新购物车的购物项数和总计金额
			c, err := model.GetCartByID(cit.CartID)
			err = model.UpdateCountAndAmountOfCart(c)
			if err != nil {
				log.Println(err)
				util.ResponseWriteJsonOfInsideServer(w)
				return
			}
		} else {
			// 购物车已经存在该产品的购物项，该购物项的产品数量增加
			citem.Count += cit.Count
			// 数据库更新购物项的产品数量和小计金额
			err = model.UpdateProductCountOfCartItem(citem)
			if err != nil {
				log.Println(err)
				util.ResponseWriteJsonOfInsideServer(w)
				return
			}
			// 数据库更新购物车的购物项数和总计金额
			c, err := model.GetCartByID(cit.CartID)
			err = model.UpdateCountAndAmountOfCart(c)
			if err != nil {
				log.Println(err)
				util.ResponseWriteJsonOfInsideServer(w)
				return
			}
		}

		// 返回成功的JSON响应
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库添加购物项成功",
		})

	case http.MethodPut:
		// 更新购物项

		// 维护购物项结构体的产品字段
		strProductID := strconv.Itoa(cit.ProductID)
		product, err := model.GetProduct(strProductID)
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		cit.Product = product

		// 数据库更新购物项的产品数量
		if err = model.UpdateProductCountOfCartItem(&cit); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		// 数据库更新购物车的购物项数和总计金额
		c, err := model.GetCartByID(cit.CartID)
		err = model.UpdateCountAndAmountOfCart(c)
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		// 返回成功的JSON响应
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库更新购物项成功",
		})

	case http.MethodDelete:
		// 删除购物项

		// 维护购物项结构体的产品字段
		strProductID := strconv.Itoa(cit.ProductID)
		product, err := model.GetProduct(strProductID)
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		cit.Product = product

		// 数据库删除购物项
		if err = cit.Delete(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		// 数据库更新购物车的购物项数和总计金额
		c, err := model.GetCartByID(cit.CartID)
		err = model.UpdateCountAndAmountOfCart(c)
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		// 返回成功的JSON响应
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库删除购物项成功",
		})
	}
}
