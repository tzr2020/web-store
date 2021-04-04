package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"web-store/model"
	"web-store/util"
)

func registerAPIOrderRouters() {
	http.HandleFunc("/api/orders", orders)
	http.HandleFunc("/api/order", order)
	http.HandleFunc("/api/order/payment-type-list", orderPaymentTypeList)
	http.HandleFunc("/api/order/payment-type-list2", orderPaymentTypeList2)
	http.HandleFunc("/api/order/payment-type", orderPaymentType)
	http.HandleFunc("/api/order/status-list", orderStatusList)
	http.HandleFunc("/api/order/status-list2", orderStatusList2)
	http.HandleFunc("/api/order/status", orderStatus)
	http.HandleFunc("/api/order/address-list", orderAddressList)
	http.HandleFunc("/api/order/address", orderAddress)
	http.HandleFunc("/api/order/item-list", orderItemList)
	http.HandleFunc("/api/order/item", orderItem)

}

// orders 获取订单列表，根据当前页的页码和每页记录条数
func orders(w http.ResponseWriter, r *http.Request) {
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
		count, err := util.GetJsonDataCount("orders")
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		// 查询数据库，获取列表，根据当前页的页码和每页记录条数
		list, err := model.GetOrderPage(intPageNo, intPageSize)
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		// 将JSON结构编码为JSON格式数据后写入响应
		j := util.Json{
			Code:  200,
			Msg:   "获取订单列表成功",
			Count: count,
			Data:  list,
		}
		util.ResponseWriteJson(w, j)
	}
}

func order(w http.ResponseWriter, r *http.Request) {
	// 解码包体JSON数据到结构
	dec := json.NewDecoder(r.Body)
	o := model.Order{}
	err := dec.Decode(&o)
	if err != nil {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}

	switch r.Method {
	case http.MethodPost:
		// 添加订单

		// 数据库添加订单
		if err = o.Add2(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库添加订单成功",
		})
	case http.MethodPut:
		// 更新订单

		// 数据库更新订单
		if err = o.Update(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库更新订单成功",
		})
	case http.MethodDelete:
		// 删除订单

		// 数据库删除订单
		if err = o.Delete(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库删除订单成功",
		})
	}
}

// orderPaymentTypeList 获取订单支付类型列表，根据当前页的页码和每页记录条数
func orderPaymentTypeList(w http.ResponseWriter, r *http.Request) {
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
		count, err := util.GetJsonDataCount("order_payment_type")
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		// 查询数据库，获取列表，根据当前页的页码和每页记录条数
		list, err := model.GetOrderPaymentTypePage(intPageNo, intPageSize)
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		// 将JSON结构编码为JSON格式数据后写入响应
		j := util.Json{
			Code:  200,
			Msg:   "获取订单支付类型列表成功",
			Count: count,
			Data:  list,
		}
		util.ResponseWriteJson(w, j)
	}
}

// orderPaymentTypeList2 获取订单支付类型列表
func orderPaymentTypeList2(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// 获取订单支付类型列表

		// 查询数据库，获取订单支付类型列表
		list, err := model.GetOrderPaymentTypes()
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		// 返回成功的JSON数据
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "获取订单支付类型列表成功",
			Data: list,
		})
	}
}

func orderPaymentType(w http.ResponseWriter, r *http.Request) {
	// 解码包体JSON数据到结构
	dec := json.NewDecoder(r.Body)
	opt := model.OrderPaymentType{}
	err := dec.Decode(&opt)
	if err != nil {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}

	switch r.Method {
	case http.MethodPost:
		// 添加订单支付类型

		// 数据库添加订单支付类型
		if err = opt.Add(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库添加订单支付类型成功",
		})
	case http.MethodPut:
		// 更新订单支付类型

		// 数据库更新订单支付类型
		if err = opt.Update(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库更新订单支付类型成功",
		})
	case http.MethodDelete:
		// 删除订单支付类型

		// 数据库删除订单支付类型
		if err = opt.Delete(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库删除订单支付类型成功",
		})
	}
}

// orderStatusList 获取订单状态字典列表，根据当前页的页码和每页记录条数
func orderStatusList(w http.ResponseWriter, r *http.Request) {
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
		count, err := util.GetJsonDataCount("order_status")
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		// 查询数据库，获取列表，根据当前页的页码和每页记录条数
		list, err := model.GetOrderStatusPage(intPageNo, intPageSize)
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		// 将JSON结构编码为JSON格式数据后写入响应
		j := util.Json{
			Code:  200,
			Msg:   "获取订单状态字典列表成功",
			Count: count,
			Data:  list,
		}
		util.ResponseWriteJson(w, j)
	}
}

// orderStatusList2 获取订单状态字典列表
func orderStatusList2(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// 获取订单状态字典列表

		// 查询数据库，获取订单状态字典列表
		list, err := model.GetOrderStatus()
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		// 返回成功的JSON数据
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "获取订单状态字典列表成功",
			Data: list,
		})
	}
}

func orderStatus(w http.ResponseWriter, r *http.Request) {
	// 解码包体JSON数据到结构
	dec := json.NewDecoder(r.Body)
	osobj := model.OrderStatus{}
	err := dec.Decode(&osobj)
	if err != nil {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}

	switch r.Method {
	case http.MethodPost:
		// 添加订单状态字典

		// 数据库添加订单状态字典
		if err = osobj.Add(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库添加订单状态字典成功",
		})
	case http.MethodPut:
		// 更新订单状态字典

		// 数据库更新订单状态字典
		if err = osobj.Update(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库更新订单状态字典成功",
		})
	case http.MethodDelete:
		// 删除订单状态字典

		// 数据库删除订单状态字典
		if err = osobj.Delete(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库删除订单状态字典成功",
		})
	}
}

// orderAddressList 获取订单地址列表，根据当前页的页码和每页记录条数
func orderAddressList(w http.ResponseWriter, r *http.Request) {
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
		count, err := util.GetJsonDataCount("order_addresses")
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		// 查询数据库，获取列表，根据当前页的页码和每页记录条数
		list, err := model.GetOrderAddressPage(intPageNo, intPageSize)
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		// 将JSON结构编码为JSON格式数据后写入响应
		j := util.Json{
			Code:  200,
			Msg:   "获取订单地址列表成功",
			Count: count,
			Data:  list,
		}
		util.ResponseWriteJson(w, j)
	}
}

func orderAddress(w http.ResponseWriter, r *http.Request) {
	// 解码包体JSON数据到结构
	dec := json.NewDecoder(r.Body)
	oa := model.OrderAddress{}
	err := dec.Decode(&oa)
	if err != nil {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}

	switch r.Method {
	case http.MethodPost:
		// 添加订单地址

		// 数据库添加订单地址
		if err = oa.Add(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库添加订单地址成功",
		})
	case http.MethodPut:
		// 更新订单地址

		// 数据库更新订单地址
		if err = oa.Update(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库更新订单地址成功",
		})
	case http.MethodDelete:
		// 删除订单地址

		// 数据库删除订单地址
		if err = oa.Delete(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库删除订单地址成功",
		})
	}
}

// orderItemList 获取订单项列表，根据当前页的页码和每页记录条数
func orderItemList(w http.ResponseWriter, r *http.Request) {
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
		count, err := util.GetJsonDataCount("order_items")
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		// 查询数据库，获取列表，根据当前页的页码和每页记录条数
		list, err := model.GetOrderItemPage(intPageNo, intPageSize)
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		// 将JSON结构编码为JSON格式数据后写入响应
		j := util.Json{
			Code:  200,
			Msg:   "获取订单项列表成功",
			Count: count,
			Data:  list,
		}
		util.ResponseWriteJson(w, j)
	}
}

func orderItem(w http.ResponseWriter, r *http.Request) {
	// 解码包体JSON数据到结构
	dec := json.NewDecoder(r.Body)
	oit := model.OrderItem{}
	err := dec.Decode(&oit)
	if err != nil {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}

	switch r.Method {
	case http.MethodPost:
		// 添加订单项

		// 数据库添加订单项
		if err = oit.Add(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库添加订单项成功",
		})
	case http.MethodPut:
		// 更新订单项

		// 数据库更新订单项
		if err = oit.Update(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库更新订单项成功",
		})
	case http.MethodDelete:
		// 删除订单项

		// 数据库删除订单项
		if err = oit.Delete(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库删除订单项成功",
		})
	}
}
