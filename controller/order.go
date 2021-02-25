package controller

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"web-store/model"
	"web-store/util"
)

func registerOrderRoutes() {
	http.HandleFunc("/writeOrder", WriteOrder)
}

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
		log.Println("数据类型转换发生错误:", err)
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

	addresses, err := model.GetAddressByUserID(sess.UserID)
	if err != nil && err != sql.ErrNoRows {
		log.Println("从数据库获取收货地址发生错误:", err)
		http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
		return
	}
	if addresses == nil {
		log.Println("用户的购物车为空")
		w.Write([]byte("你还没有添加收货地址，请先去添加收货地址。"))
		return
	}
	address := addresses[0]

	shipFee := 11.00
	order := &model.Order{
		Payment:     cart.TotalAmount + shipFee,
		PaymentType: intPaymentType,
		ShipFee:     shipFee,
	}

	sess.Cart = cart
	sess.Order = order
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
