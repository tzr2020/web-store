package controller

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
	"web-store/model"
	"web-store/util"
)

func registerOrderRoutes() {
	http.HandleFunc("/writeOrder", WriteOrder)
	http.HandleFunc("/commitOrder", CommitOrder)
	http.HandleFunc("/myOrder", MyOrder)
}

// WriteOrder 用户填写订单
func WriteOrder(w http.ResponseWriter, r *http.Request) {
	ok, sess := IsLogin(r)

	if !ok {
		log.Println("用户填写订单时，没有登录账号")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("你没有登录账号。"))
		return
	}

	paymentType := r.PostFormValue("paymentType")
	if paymentType == "" {
		paymentType = "1"
	}
	intPaymentType, err := strconv.Atoi(paymentType)
	if err != nil {
		log.Println("订单的支付方式字段类型转换发生错误:", err)
		http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
		return
	}

	cart, err := model.GetCartByUserID(sess.UserID)
	if err != nil && err != sql.ErrNoRows {
		log.Println("从数据库获取购物车发生错误:", err)
		http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
		return
	}
	if err == sql.ErrNoRows {
		log.Println("用户的购物车为空，无法去填写订单")
		w.Write([]byte("你的购物车还没有产品，请先去添加产品。"))
		return
	}

	address := &model.Address{}
	addresses, err := model.GetAddressByUserID(sess.UserID)
	if err != nil && err != sql.ErrNoRows {
		log.Println("从数据库获取收货地址发生错误:", err)
		http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
		return
	}
	if addresses != nil {
		address = addresses[0]
	}

	shipFee := 11.00
	order := &model.Order{
		Payment:     cart.TotalAmount + shipFee,
		PaymentType: intPaymentType,
		ShipFee:     shipFee,
	}

	orderPaymentTypes, err := model.GetOrderPaymentTypes()
	if err != nil {
		log.Println("从数据库获取所有订单支付方式发生错误:", err)
		http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
		return
	}

	sess.Cart = cart
	sess.Order = order
	sess.OrderPaymentTypes = orderPaymentTypes
	sess.Address = address

	// 解析模板文件，并执行模板，生成包含动态数据的HTML文档，返回给浏览器
	t, err := template.ParseFiles("./view/template/layout.html", "./view/template/order-write.html")
	if err != nil {
		log.Printf("解析模板文件发生错误：%v\n", err)
		http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
	} else {
		t.ExecuteTemplate(w, "layout", sess)
	}

}

// CommitOrder 用户提交订单
func CommitOrder(w http.ResponseWriter, r *http.Request) {
	ok, sess := IsLogin(r)

	if !ok {
		log.Println("用户提交订单时，没有登录账号")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("你没有登录账号。"))
		return
	}

	name := r.PostFormValue("name")
	tel := r.PostFormValue("tel")
	province := r.PostFormValue("province")
	city := r.PostFormValue("city")
	area := r.PostFormValue("area")
	street := r.PostFormValue("street")
	code := r.PostFormValue("code")
	paymentType := r.PostFormValue("paymentType")

	cart, err := model.GetCartByUserID(sess.UserID)
	if err != nil && err != sql.ErrNoRows {
		log.Println("从数据库获取购物车发生错误:", err)
		http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
		return
	}
	if err == sql.ErrNoRows {
		log.Println("用户的购物车为空，无法去提交订单")
		w.Write([]byte("你的购物车还没有产品，请先去添加产品。"))
		return
	}

	orderID := util.CreateUUID()
	intPaymentType, err := strconv.Atoi(paymentType)
	if err != nil {
		log.Println("订单的支付方式字段类型转换发生错误:", err)
		http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
		return
	}
	shipFee := 11.00
	createTime := time.Now().Format("2006-01-02 15:04:05")

	// 数据库添加订单
	order := &model.Order{
		ID:          orderID,
		UID:         sess.UserID,
		TotalCount:  cart.TotalCount,
		TotalAmount: cart.TotalAmount,
		Payment:     cart.TotalAmount + shipFee,
		PaymentType: intPaymentType,
		ShipFee:     shipFee,
		CreateTime:  createTime,
	}
	if err = order.Add(); err != nil {
		log.Println("数据库添加订单发生错误:", err)
		http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
		return
	}

	// 数据库添加订单项
	for _, v := range cart.CartItems {
		orderItem := &model.OrderItem{
			OrderID:   orderID,
			ProductID: v.Product.ID,
			Count:     v.Count,
			Amount:    v.Amount,
		}
		if err = orderItem.Add(); err != nil {
			log.Println("数据库添加订单项发生错误:", err)
			http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
			return
		}

		// 数据库更新产品的库存和销量
		product := v.Product
		product.Stock = product.Stock - v.Count
		product.Sales = product.Sales + v.Count
		if err := product.UpdateStockAndSales(); err != nil {
			log.Println("数据库更新产品的库存和销量发生错误:", err)
			http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
			return
		}
	}

	// 数据库添加订单地址
	orderAddress := &model.OrderAddress{
		OrderID:  orderID,
		Name:     name,
		Tel:      tel,
		Province: province,
		City:     city,
		Area:     area,
		Strees:   street,
		Code:     code,
	}
	if err = orderAddress.Add(); err != nil {
		log.Println("数据库添加订单地址发生错误:", err)
		http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
		return
	}

	// 数据库删除购物车
	if err = model.DeleteCart(cart.CartID); err != nil {
		log.Println("数据库删除购物车发生错误:", err)
		http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("成功提交订单。"))
}

// MyOrder 用户查看订单，返回用户订单页面
func MyOrder(w http.ResponseWriter, r *http.Request) {
	ok, sess := IsLogin(r)

	if !ok {
		log.Println("用户查看订单时，没有登录账号")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("你没有登录账号。"))
		return
	}

	orders, err := model.GetOrdersByUserID(sess.UserID)
	if err != nil {
		log.Println("数据库获取某用户的所有订单发生错误:", err)
		http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
		return
	}

	sess.Orders = orders

	// 解析模板文件，并执行模板，生成包含动态数据的HTML文档，返回给浏览器
	funcMap := template.FuncMap{"OrderPaymentTypeCodeToText": model.OrderPaymentTypeCodeToText,
		"OrderStatusCodeToText": model.OrderStatusCodeToText} // 包含自定义的模板函数
	t := template.New("layout").Funcs(funcMap) // 创建模板并绑定FuncMap
	t, err = t.ParseFiles("./view/template/layout.html", "./view/template/order.html")
	if err != nil {
		log.Printf("解析模板文件发生错误：%v\n", err)
		http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
	} else {
		t.ExecuteTemplate(w, "layout", sess)
	}
}
