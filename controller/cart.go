package controller

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"web-store/model"
	"web-store/util"
)

func regsiterCartRoutes() {
	http.HandleFunc("/addToCart", AddToCart)
	http.HandleFunc("/getCartInfo", GetCartInfo)
	http.HandleFunc("/deleteCart", DeleteCart)
}

// AddToCart 将产品添加到购物车
// 整体流程：用户是否已经登录->数据库是否存在购物车->数据库是否存在购物项
func AddToCart(w http.ResponseWriter, r *http.Request) {
	ok, sess := IsLogin(r)

	if ok {
		// 用户已经登录

		pid := r.FormValue("productID")
		uid := sess.UserID

		// 从数据库获取产品
		p, err := model.GetProduct(pid)
		if err != nil {
			log.Println("从数据库获取产品发生错误: ", err)
			http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
			return
		}

		// 从数据库获取购物车
		c, err := model.GetCartByUserID(uid)

		// test
		// fmt.Println("购物车: ", c)
		// if c != nil {
		// 	for k, v := range c.CartItems {
		// 		fmt.Printf("第%v个购物项的信息: %v，产品信息：%v\n", k+1, v, v.Product)
		// 	}
		// }

		if c != nil {
			// 用户在数据库里已有购物车

			// 从数据库获取某产品购物项
			cItem, err := model.GetCartItemByCartIDAndProductID(c.CartID, pid)
			// tset
			// fmt.Println("购物项: ", cItem)
			if err != nil && err != sql.ErrNoRows {
				log.Println("从数据库获取购物项发生错误: ", err)
				http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
				return
			}

			if cItem != nil {
				// 购物车已有该产品对应的购物项，该产品的数量加1
				for _, v := range c.CartItems {
					if v.Product.ID == cItem.Product.ID {
						v.Count += 1
						// test
						// fmt.Println(v.Count, v.Product.ID, c.CartID)
						err := model.UpdateProductCountOfCartItem(v)
						if err != nil {
							log.Println("数据库更新购物项的产品数量发生错误: ", err)
							http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
							return
						}
					}
				}
			} else {
				// test
				// fmt.Println("产品: ", p)
				// 购物车还没有该产品对应的购物项，创建购物项添加到数据库，并维护购物车的购物项字段
				cItem := &model.CartItem{
					CartID:  c.CartID,
					Product: p,
					Count:   1,
				}
				err := model.AddCartItem(cItem)
				if err != nil {
					log.Println("新增购物项到数据库发生错误: ", err)
					http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
					return
				}
				c.CartItems = append(c.CartItems, cItem)
			}

			// test
			// fmt.Println("更新前的购物车：", c)
			// if c != nil {
			// 	for k, v := range c.CartItems {
			// 		fmt.Printf("第%v个购物项的信息: %v，产品信息：%v\n", k+1, v, v.Product)
			// 	}
			// }

			// 无论有没有购物车有没有该产品对应的购物项，最后都需要更新购物车
			err = model.UpdateCountAndAmountOfCart(c)
			if err != nil {
				log.Println("数据库更新购物车的购物项数和总计金额发生错误: ", err)
				http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
				return
			}

		} else {
			// 用户在数据库里还没购物车，数据库新增购物车
			cid := util.CreateUUID()

			cItem := &model.CartItem{
				CartID:  cid,
				Product: p,
				Count:   1,
			}

			var cItems []*model.CartItem
			cItems = append(cItems, cItem)

			cart := &model.Cart{
				CartID:    cid,
				UserID:    uid,
				CartItems: cItems,
			}

			err := model.AddCart(cart)
			if err != nil {
				log.Println("新增购物车到数据库发生错误: ", err)
				http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
				return
			}
		}

		w.Write([]byte("产品: " + p.Name + " 已加入购物车"))

	} else {
		// 用户还没登录
		w.Write([]byte("请先登录账号，再将产品加入购物车"))
	}
}

// GetCartInfo 返回购物车内容页面
func GetCartInfo(w http.ResponseWriter, r *http.Request) {
	ok, sess := IsLogin(r)

	if !ok {
		// 解析模板文件，执行模板结合动态数据，生成最终HTML文档，传递给ResponseWriter响应客户端
		t, err := template.ParseFiles("./view/template/layout.html", "./view/template/cart-info.html")
		if err != nil {
			log.Printf("解析模板文件发生错误：%v\n", err)
			http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
		} else {
			t.ExecuteTemplate(w, "layout", nil)
		}
		return
	}

	cart, err := model.GetCartByUserID(sess.UserID)
	if err != nil {
		log.Println("从数据库获取购物车发生错误: ", err)
	}

	if cart != nil {
		sess.Cart = cart
	}

	// 解析模板文件，执行模板结合动态数据，生成最终HTML文档，传递给ResponseWriter响应客户端
	t, err := template.ParseFiles("./view/template/layout.html", "./view/template/cart-info.html")
	if err != nil {
		log.Printf("解析模板文件发生错误：%v\n", err)
		http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
	} else {
		t.ExecuteTemplate(w, "layout", sess)
	}
}

// DeleteCart 清空购物车，返回购物车内容页面
func DeleteCart(w http.ResponseWriter, r *http.Request) {
	cartID := r.FormValue("cartID")

	err := model.DeleteCart(cartID)
	if err != nil {
		log.Println("从数据库删除购物车发生错误:", err)
		http.Error(w, util.ErrServerInside.Error(), http.StatusInternalServerError)
		return
	}

	GetCartInfo(w, r)
}
