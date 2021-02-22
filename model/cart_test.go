package model

import (
	"fmt"
	"log"
	"testing"
	"web-store/util"
)

func TestCart(t *testing.T) {
	// t.Run("测试添加购物车", testAddCart)
	// t.Run("测试获取购物车，根据用户id", testGetCartByUserID)
}

func testAddCart(t *testing.T) {
	fmt.Println("测试添加购物车")

	cartID := util.CreateUUID()

	p, err := GetProduct("1")
	if err != nil {
		log.Println(err)
		return
	}
	p2, err := GetProduct("2")
	if err != nil {
		log.Println(err)
		return
	}

	cItem := &CartItem{
		CartID:  cartID,
		Product: p,
		Count:   10,
	}
	cItem2 := &CartItem{
		CartID:  cartID,
		Product: p2,
		Count:   10,
	}

	var cartItems []*CartItem
	cartItems = append(cartItems, cItem)
	cartItems = append(cartItems, cItem2)

	cart := &Cart{
		CartID:    cartID,
		UserID:    1,
		CartItems: cartItems,
	}

	err = AddCart(cart)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("添加购物车成功")
}

func testGetCartByUserID(t *testing.T) {
	fmt.Println("测试获取购物车，根据用户id")

	cart, err := GetCartByUserID(1)
	if err != nil {
		return
	}

	fmt.Printf("购物车信息: %v\n", cart)
	for k, v := range cart.CartItems {
		fmt.Printf("第%v个购物项信息: %v\n", k+1, v)
	}
}
