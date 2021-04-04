package controller

import (
	"database/sql"
	"log"
	"net/http"
	"web-store/util"
)

func registerManageRoutes() {
	http.HandleFunc("/manage/index", index)
	http.HandleFunc("/manage/users", usersPage)
	http.HandleFunc("/manage/user/addresses", usersAddressesPage)
	http.HandleFunc("/manage/products", productsPage)
	http.HandleFunc("/manage/carts", cartsPage)
	http.HandleFunc("/manage/cartitems", cartitemsPage)
	http.HandleFunc("/manage/categories", categoriesPage)
	http.HandleFunc("/manage/index/hot-products", indexHotProductsPage)
	http.HandleFunc("/manage/index/new-products", indexNewProductsPage)
	http.HandleFunc("/manage/index/recom-products", indexRecomProductsPage)
	http.HandleFunc("/manage/nav-products", navProductsPage)
	http.HandleFunc("/manage/orders", ordersPage)
	http.HandleFunc("/manage/order/payment-type-list", orderPaymentTypeListPage)
	http.HandleFunc("/manage/order/status-list", orderStatusListPage)
	http.HandleFunc("/manage/order/address-list", orderAddressListPage)
	http.HandleFunc("/manage/order/item-list", orderItemListPage)
	http.HandleFunc("/manage/user/session-list", userSessionListPage)
	http.HandleFunc("/manage/admin/session-list", adminSessionListPage)
	http.HandleFunc("/manage/admin-list", adminListPage)
}

// index 返回后台管理系统的首页
func index(w http.ResponseWriter, r *http.Request) {
	// 判断管理员是否已经登录
	ok, _, err := AdminIsLoginByCookie(r)
	if err != nil && err != http.ErrNoCookie && err != sql.ErrNoRows {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}
	if !ok {
		// 重定向到管理员登录页面
		w.Header().Set("Location", "/api/admin/login")
		w.WriteHeader(302)
		return
	}
	// 返回后台管理系统首页
	util.ExecuteTowTpl(w, [2]string{
		"./view/template/manage/manage-layout.html",
		"./view/template/manage/manage-index.html",
	})
}

// usersPage 返回后台管理系统的会员用户列表页面
func usersPage(w http.ResponseWriter, r *http.Request) {
	// 判断管理员是否已经登录
	ok, _, err := AdminIsLoginByCookie(r)
	if err != nil && err != http.ErrNoCookie && err != sql.ErrNoRows {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}
	if !ok {
		// 重定向到管理员登录页面
		w.Header().Set("Location", "/api/admin/login")
		w.WriteHeader(302)
		return
	}
	// 返回后台管理系统的会员用户列表页面
	util.ExecuteTowTpl(w, [2]string{
		"./view/template/manage/manage-layout.html",
		"./view/template/manage/users.html",
	})
}

// usersAddressesPage 返回后台管理系统的会员用户地址列表页面
func usersAddressesPage(w http.ResponseWriter, r *http.Request) {
	// 判断管理员是否已经登录
	ok, _, err := AdminIsLoginByCookie(r)
	if err != nil && err != http.ErrNoCookie && err != sql.ErrNoRows {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}
	if !ok {
		// 重定向到管理员登录页面
		w.Header().Set("Location", "/api/admin/login")
		w.WriteHeader(302)
		return
	}
	// 返回后台管理系统的会员用户地址列表页面
	util.ExecuteTowTpl(w, [2]string{
		"./view/template/manage/manage-layout.html",
		"./view/template/manage/user-addresses.html",
	})
}

// productsPage 返回后台管理系统的产品列表页面
func productsPage(w http.ResponseWriter, r *http.Request) {
	// 判断管理员是否已经登录
	ok, _, err := AdminIsLoginByCookie(r)
	if err != nil && err != http.ErrNoCookie && err != sql.ErrNoRows {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}
	if !ok {
		// 重定向到管理员登录页面
		w.Header().Set("Location", "/api/admin/login")
		w.WriteHeader(302)
		return
	}
	// 返回后台管理系统的产品列表页面
	util.ExecuteTowTpl(w, [2]string{
		"./view/template/manage/manage-layout.html",
		"./view/template/manage/products.html",
	})
}

// cartsPage 返回后台管理系统的购物车列表页面
func cartsPage(w http.ResponseWriter, r *http.Request) {
	// 判断管理员是否已经登录
	ok, _, err := AdminIsLoginByCookie(r)
	if err != nil && err != http.ErrNoCookie && err != sql.ErrNoRows {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}
	if !ok {
		// 重定向到管理员登录页面
		w.Header().Set("Location", "/api/admin/login")
		w.WriteHeader(302)
		return
	}
	// 返回后台管理系统的产品列表页面
	util.ExecuteTowTpl(w, [2]string{
		"./view/template/manage/manage-layout.html",
		"./view/template/manage/carts.html",
	})
}

// cartitemsPage 返回后台管理系统的购物项列表页面
func cartitemsPage(w http.ResponseWriter, r *http.Request) {
	// 判断管理员是否已经登录
	ok, _, err := AdminIsLoginByCookie(r)
	if err != nil && err != http.ErrNoCookie && err != sql.ErrNoRows {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}
	if !ok {
		// 重定向到管理员登录页面
		w.Header().Set("Location", "/api/admin/login")
		w.WriteHeader(302)
		return
	}
	// 返回后台管理系统的购物项列表页面
	util.ExecuteTowTpl(w, [2]string{
		"./view/template/manage/manage-layout.html",
		"./view/template/manage/cartitems.html",
	})
}

// categoriesPage 返回后台管理系统的产品类别列表页面
func categoriesPage(w http.ResponseWriter, r *http.Request) {
	// 判断管理员是否已经登录
	ok, _, err := AdminIsLoginByCookie(r)
	if err != nil && err != http.ErrNoCookie && err != sql.ErrNoRows {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}
	if !ok {
		// 重定向到管理员登录页面
		w.Header().Set("Location", "/api/admin/login")
		w.WriteHeader(302)
		return
	}
	// 返回后台管理系统的产品类别列表页面
	util.ExecuteTowTpl(w, [2]string{
		"./view/template/manage/manage-layout.html",
		"./view/template/manage/categories.html",
	})
}

// indexHotProductsPage 返回后台管理系统的首页热卖产品列表页面
func indexHotProductsPage(w http.ResponseWriter, r *http.Request) {
	// 判断管理员是否已经登录
	ok, _, err := AdminIsLoginByCookie(r)
	if err != nil && err != http.ErrNoCookie && err != sql.ErrNoRows {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}
	if !ok {
		// 重定向到管理员登录页面
		w.Header().Set("Location", "/api/admin/login")
		w.WriteHeader(302)
		return
	}
	// 返回后台管理系统的首页热卖产品列表页面
	util.ExecuteTowTpl(w, [2]string{
		"./view/template/manage/manage-layout.html",
		"./view/template/manage/index-hot-products.html",
	})
}

// indexNewProductsPage 返回后台管理系统的首页最新产品列表页面
func indexNewProductsPage(w http.ResponseWriter, r *http.Request) {
	// 判断管理员是否已经登录
	ok, _, err := AdminIsLoginByCookie(r)
	if err != nil && err != http.ErrNoCookie && err != sql.ErrNoRows {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}
	if !ok {
		// 重定向到管理员登录页面
		w.Header().Set("Location", "/api/admin/login")
		w.WriteHeader(302)
		return
	}
	// 返回后台管理系统的首页最新产品列表页面
	util.ExecuteTowTpl(w, [2]string{
		"./view/template/manage/manage-layout.html",
		"./view/template/manage/index-new-products.html",
	})
}

// indexRecomProductsPage 返回后台管理系统的首页推荐产品列表页面
func indexRecomProductsPage(w http.ResponseWriter, r *http.Request) {
	// 判断管理员是否已经登录
	ok, _, err := AdminIsLoginByCookie(r)
	if err != nil && err != http.ErrNoCookie && err != sql.ErrNoRows {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}
	if !ok {
		// 重定向到管理员登录页面
		w.Header().Set("Location", "/api/admin/login")
		w.WriteHeader(302)
		return
	}
	// 返回后台管理系统的首页推荐产品列表页面
	util.ExecuteTowTpl(w, [2]string{
		"./view/template/manage/manage-layout.html",
		"./view/template/manage/index-recom-products.html",
	})
}

// navProductsPage 返回后台管理系统的导航栏产品列表页面
func navProductsPage(w http.ResponseWriter, r *http.Request) {
	// 判断管理员是否已经登录
	ok, _, err := AdminIsLoginByCookie(r)
	if err != nil && err != http.ErrNoCookie && err != sql.ErrNoRows {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}
	if !ok {
		// 重定向到管理员登录页面
		w.Header().Set("Location", "/api/admin/login")
		w.WriteHeader(302)
		return
	}
	// 返回后台管理系统的导航栏产品列表页面
	util.ExecuteTowTpl(w, [2]string{
		"./view/template/manage/manage-layout.html",
		"./view/template/manage/nav-products.html",
	})
}

// ordersPage 返回后台管理系统的订单列表页面
func ordersPage(w http.ResponseWriter, r *http.Request) {
	// 判断管理员是否已经登录
	ok, _, err := AdminIsLoginByCookie(r)
	if err != nil && err != http.ErrNoCookie && err != sql.ErrNoRows {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}
	if !ok {
		// 重定向到管理员登录页面
		w.Header().Set("Location", "/api/admin/login")
		w.WriteHeader(302)
		return
	}
	// 返回后台管理系统的订单列表页面
	util.ExecuteTowTpl(w, [2]string{
		"./view/template/manage/manage-layout.html",
		"./view/template/manage/orders.html",
	})
}

// orderPaymentTypeListPage 返回后台管理系统的订单支付类型列表页面
func orderPaymentTypeListPage(w http.ResponseWriter, r *http.Request) {
	// 判断管理员是否已经登录
	ok, _, err := AdminIsLoginByCookie(r)
	if err != nil && err != http.ErrNoCookie && err != sql.ErrNoRows {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}
	if !ok {
		// 重定向到管理员登录页面
		w.Header().Set("Location", "/api/admin/login")
		w.WriteHeader(302)
		return
	}
	// 返回后台管理系统的订单支付类型列表页面
	util.ExecuteTowTpl(w, [2]string{
		"./view/template/manage/manage-layout.html",
		"./view/template/manage/order-payment-type.html",
	})
}

// orderStatusListPage 返回后台管理系统的订单状态字典列表页面
func orderStatusListPage(w http.ResponseWriter, r *http.Request) {
	// 判断管理员是否已经登录
	ok, _, err := AdminIsLoginByCookie(r)
	if err != nil && err != http.ErrNoCookie && err != sql.ErrNoRows {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}
	if !ok {
		// 重定向到管理员登录页面
		w.Header().Set("Location", "/api/admin/login")
		w.WriteHeader(302)
		return
	}
	// 返回后台管理系统的订单状态字典列表页面
	util.ExecuteTowTpl(w, [2]string{
		"./view/template/manage/manage-layout.html",
		"./view/template/manage/order-status.html",
	})
}

// orderAddressListPage 返回后台管理系统的订单地址列表页面
func orderAddressListPage(w http.ResponseWriter, r *http.Request) {
	// 判断管理员是否已经登录
	ok, _, err := AdminIsLoginByCookie(r)
	if err != nil && err != http.ErrNoCookie && err != sql.ErrNoRows {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}
	if !ok {
		// 重定向到管理员登录页面
		w.Header().Set("Location", "/api/admin/login")
		w.WriteHeader(302)
		return
	}
	// 返回后台管理系统的订单地址列表页面
	util.ExecuteTowTpl(w, [2]string{
		"./view/template/manage/manage-layout.html",
		"./view/template/manage/order-address.html",
	})
}

// orderItemListPage 返回后台管理系统的订单项列表页面
func orderItemListPage(w http.ResponseWriter, r *http.Request) {
	// 判断管理员是否已经登录
	ok, _, err := AdminIsLoginByCookie(r)
	if err != nil && err != http.ErrNoCookie && err != sql.ErrNoRows {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}
	if !ok {
		// 重定向到管理员登录页面
		w.Header().Set("Location", "/api/admin/login")
		w.WriteHeader(302)
		return
	}
	// 返回后台管理系统的订单项列表页面
	util.ExecuteTowTpl(w, [2]string{
		"./view/template/manage/manage-layout.html",
		"./view/template/manage/order-item.html",
	})
}

// userSessionListPage 返回后台管理系统的用户Session列表页面
func userSessionListPage(w http.ResponseWriter, r *http.Request) {
	// 判断管理员是否已经登录
	ok, _, err := AdminIsLoginByCookie(r)
	if err != nil && err != http.ErrNoCookie && err != sql.ErrNoRows {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}
	if !ok {
		// 重定向到管理员登录页面
		w.Header().Set("Location", "/api/admin/login")
		w.WriteHeader(302)
		return
	}
	// 返回后台管理系统的用户Session列表页面
	util.ExecuteTowTpl(w, [2]string{
		"./view/template/manage/manage-layout.html",
		"./view/template/manage/user-session.html",
	})
}

// adminSessionListPage 返回后台管理系统的管理员Session列表页面
func adminSessionListPage(w http.ResponseWriter, r *http.Request) {
	// 判断管理员是否已经登录
	ok, _, err := AdminIsLoginByCookie(r)
	if err != nil && err != http.ErrNoCookie && err != sql.ErrNoRows {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}
	if !ok {
		// 重定向到管理员登录页面
		w.Header().Set("Location", "/api/admin/login")
		w.WriteHeader(302)
		return
	}
	// 返回后台管理系统的管理员Session列表页面
	util.ExecuteTowTpl(w, [2]string{
		"./view/template/manage/manage-layout.html",
		"./view/template/manage/admin-session.html",
	})
}

// adminListPage 返回后台管理系统的管理员列表页面
func adminListPage(w http.ResponseWriter, r *http.Request) {
	// 判断管理员是否已经登录
	ok, _, err := AdminIsLoginByCookie(r)
	if err != nil && err != http.ErrNoCookie && err != sql.ErrNoRows {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}
	if !ok {
		// 重定向到管理员登录页面
		w.Header().Set("Location", "/api/admin/login")
		w.WriteHeader(302)
		return
	}
	// 返回后台管理系统的管理员列表页面
	util.ExecuteTowTpl(w, [2]string{
		"./view/template/manage/manage-layout.html",
		"./view/template/manage/admin.html",
	})
}
